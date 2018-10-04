package p2p

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"
	"yqx_go/young_blockchain/common"
	"yqx_go/young_blockchain/consensus"
	TxModule "yqx_go/young_blockchain/transactions"
)

//Run run test
func Run() {

	listenF := flag.Int("l", 10001, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
	mineF := flag.Bool("m", false, "target peer to dial")
	flag.Parse()
	//first
	target2 := RunServer(*listenF, *target)
	fmt.Printf("peer target: %s\n", target2)
	//second
	//target3 := RunServer(10002,target2)
	//RunServer(10002, target2)
	//third
	//RunServer(10003, target2)

	//time.Sleep(20 * time.Second)
	if *mineF == true{
		cbtx := TxModule.NewCoinbaseTX(common.ToAddress, common.GenesisCoinbaseData)
		genesis := consensus.NewGenesisBlock(cbtx)
		fmt.Printf("target: %s\n", "mined a block")
		WriteABlock(genesis)
		//time.Sleep(20 * time.Second)
	}
}

//RunServer run test
func RunServer(listentPort int, target string) (string) {
	//first
	var server *p2pObject
	server = &p2pObject{listenFport: listentPort, mutex: &sync.Mutex{}}
	ha, fullAddr, err := server.InitialAddress()
	go func() {
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("StartRunner run \"go run main.go -l %d -d %s -secio\" on a different terminal\n", listentPort+1, fullAddr)
		server.StartRunner(ha, target)
	}()
	target2 := fmt.Sprintf("%s", fullAddr)
	fmt.Printf("target: %s\n", target2)
	time.Sleep(10 * time.Second)
	return target2
}
