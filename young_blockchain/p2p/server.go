package p2p

import (
	"bufio"
	"context"
	"crypto/rand"
	"sync"

	//"crypto/sha256"
	//"encoding/hex"
	"encoding/json"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	golog "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p-host"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
	gologging "github.com/whyrusleeping/go-logging"
	"io"
	"log"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"
	BlModule "yqx_go/young_blockchain/blockchain"
)

type p2pObject struct {
	listenFport  int
	mutex  *sync.Mutex
}

var Channel_Block = make(chan BlModule.Block, 3)

func (server *p2pObject) InitialAddress() (host.Host, ma.Multiaddr, error) {

	golog.SetAllLoggers(gologging.INFO) // Change to DEBUG for extra info

	// Parse options from the command line
	listenF := server.listenFport //flag.Int("l", listentPort, "wait for incoming connections")
	secio := false //flag.Bool("secio", false, "enable secio")
	seed := int64(0) //flag.Int64("seed", 0, "set random seed for id generation")

	if listenF == 0 {
		log.Fatal("Please provide a port to bind on with -l")
	}

	// Make a host that listens on the given multiaddress
	ha, addr, err := server.makeBasicHost(listenF, secio, seed)
	if err != nil {
		log.Fatal(err)
	}

	return ha, addr, nil

}

func  (server *p2pObject) StartRunner(ha host.Host,targetAddr string) {
	target := targetAddr//flag.String("d", targetAddr, "target peer to dial")
	if target == "" {
		log.Println("listening for connections")
		// Set a stream handler on host A. /p2p/1.0.0 is
		// a user-defined protocol name.
		ha.SetStreamHandler("/p2p/1.0.0", server.handleStream)

		select {} // hang forever
		/**** This is where the listener code ends ****/
	} else {
		ha.SetStreamHandler("/p2p/1.0.0", server.handleStream)

		// The following code extracts target's peer ID from the
		// given multiaddress
		ipfsaddr, err := ma.NewMultiaddr(target)
		if err != nil {
			log.Fatalln(err)
		}

		pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
		if err != nil {
			log.Fatalln(err)
		}

		peerid, err := peer.IDB58Decode(pid)
		if err != nil {
			log.Fatalln(err)
		}

		// Decapsulate the /ipfs/<peerID> part from the target
		// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
		targetPeerAddr, _ := ma.NewMultiaddr(
			fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
		targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

		// We have a peer ID and a targetAddr so we add it to the peerstore
		// so LibP2P knows how to contact it
		ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

		log.Println("opening stream")
		// make a new stream from host B to host A
		// it should be handled on host A by the handler we set above because
		// we use the same /p2p/1.0.0 protocol
		s, err := ha.NewStream(context.Background(), peerid, "/p2p/1.0.0")
		if err != nil {
			log.Fatalln(err)
		}
		// Create a buffered stream so that read and writes are non blocking.
		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

		// Create a thread to read and write data.
		go server.writeData(rw,"StartRunner")
		go server.readData(rw,"StartRunner")

		select {} // hang forever

	}
}

// makeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress. It will use secio if secio is true.
func (server *p2pObject) makeBasicHost(listenPort int, secio bool, randseed int64) (host.Host, ma.Multiaddr, error) {

	// If the seed is zero, use real cryptographic randomness. Otherwise, use a
	// deterministic randomness source to make generated keys stay the same
	// across multiple runs
	var r io.Reader
	if randseed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(randseed))
	}

	basicHost, fullAddr, err := server.createAddress(listenPort, r)
	if err != nil {
		return nil, nil, err
	}
	//log.Printf("I am1 %s\n", fullAddr)
	log.Printf("I am2 %s\n", fullAddr)
	if secio {
		log.Printf("Now run \"go run main.go -l %d -d %s -secio\" on a different terminal\n", listenPort+1, fullAddr)
	} else {
		log.Printf("Now run \"go run main.go -l %d -d %s\" on a different terminal\n", listenPort+1, fullAddr)
	}

	return basicHost, fullAddr, nil
}

func  (server *p2pObject) createAddress(listenPort int, r io.Reader) (host.Host, ma.Multiaddr, error) {

	// Generate a key pair for this host. We will use it
	// to obtain a valid host ID.
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, nil, err
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
		libp2p.Identity(priv),
	}

	basicHost, err := libp2p.New(context.Background(), opts...)
	if err != nil {
		return nil, nil, err
	}

	// Build host multiaddress
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))

	// Now we can build a full multiaddress to reach this host
	// by encapsulating both addresses:
	addr := basicHost.Addrs()[0]
	fullAddr := addr.Encapsulate(hostAddr)

	return basicHost, fullAddr, nil
}

func  (server *p2pObject) handleStream(s net.Stream) {

	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go server.readData(rw,"handleStream")
	go server.writeData(rw,"handleStream")

	// stream 's' will stay open until you close it (or the other side closes it).
}

func  (server *p2pObject) readData(rw *bufio.ReadWriter,category string) {

	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if str == "" {
			return
		}
		if str != "\n" {
			server.mutex.Lock()
			str_to_send := strings.Replace(str, "\n", "", -1)
			fmt.Printf("read to print:%s,%d,%s,%s\n", str_to_send,server.listenFport,category,time.Now())
			//rw.WriteString(fmt.Sprintf("read to write: %d,%s,%d,%s\n", 1,str_to_send, server.listenFport,category))
			//rw.Flush()
			server.mutex.Unlock()

			//read block from ReadWriter

			/*chain := make([]BlModule.Block, 0)
			if err := json.Unmarshal([]byte(str), &chain); err != nil {
				log.Fatal(err)
			}

			mutex.Lock()
			if len(chain) > len(Blockchain) {
				Blockchain = chain
				bytes, err := json.MarshalIndent(Blockchain, "", "  ")
				if err != nil {

					log.Fatal(err)
				}
				// Green console color: 	\x1b[32m
				// Reset console color: 	\x1b[0m
				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}
			mutex.Unlock()*/
		}
	}
}

func  (server *p2pObject) writeData(rw *bufio.ReadWriter,category string) {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			server.mutex.Lock()
			//write a block
			block := &BlModule.Block{Timestamp: time.Now().Unix(), Transactions: nil, PrevBlockHash: nil,
				Hash: []byte{}, Height: 0}
			bytes, err := json.Marshal(block)
			if err != nil {
				log.Println(err)
			}
			server.mutex.Unlock()

			server.mutex.Lock()
			rw.WriteString(fmt.Sprintf("writeData 1: %s,%d,%s\n", string(bytes), server.listenFport,category))
			rw.Flush()
			server.mutex.Unlock()

		}
	}()

	time.Sleep(8 * time.Second)

	for {
		server.mutex.Lock()
		v := <-Channel_Block
		//spew.Dump(v)
		fmt.Println("receive a block:", v.Height,time.Now())
		//fmt.Printf("target555: %s,%d,%s\n", string(bytes[:]),server.listenFport,category)
		rw.WriteString(fmt.Sprintf("target666 2: %d,%d,%d,%s\n", 1,v.Height, server.listenFport,category))
		rw.Flush()
		server.mutex.Unlock()
	}

	//stdReader := bufio.NewReader(os.Stdin)
	stdReader := bufio.NewReader(strings.NewReader(string("123\n")))

	for {
		fmt.Print(">***> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		sendData = strings.Replace(sendData, "\n", "", -1)
		bpm, err := strconv.Atoi(sendData)
		if err != nil {
			log.Fatal(err)
		}
		/*newBlock := generateBlock(Blockchain[len(Blockchain)-1], bpm)

		if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
			mutex.Lock()
			Blockchain = append(Blockchain, newBlock)
			mutex.Unlock()
		}

		bytes, err := json.Marshal(Blockchain)
		if err != nil {
			log.Println(err)
		}

		spew.Dump(Blockchain)*/
		bytes := []byte("byteHello")
		server.mutex.Lock()
		fmt.Printf("target555: %s,%d,%s\n", string(bytes[:]),server.listenFport,category)
		rw.WriteString(fmt.Sprintf("writeData 2: %d,%d,%d,%s\n", 1,bpm, server.listenFport,category))
		rw.Flush()
		server.mutex.Unlock()
		time.Sleep(20 * time.Second)
	}
}

func WriteABlock(block *BlModule.Block) {
	Channel_Block <- *block
	Channel_Block <- *block
	Channel_Block <- *block
}