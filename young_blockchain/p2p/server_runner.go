package p2p

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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
	flag.Parse()

	target_to := RunServer(*listenF, *target)
	fmt.Printf("peer target: %s\n", target_to)


}
//AcceptCmd from console
func AcceptCmd(){
	for {
		stdReader := bufio.NewReader(os.Stdin)
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("AcceptCmd: %s\n", sendData)
		cmd := strings.Replace(sendData, "\n", "", -1)
		RunCmd(cmd)
	}

}
//RunCmd run cmd from console
func RunCmd(cmd string) {
	switch cmd {
	case common.Miner.String():
		cbtx := TxModule.NewCoinbaseTX(common.ToAddress, common.GenesisCoinbaseData)
		genesis := consensus.NewGenesisBlock(cbtx)
		fmt.Printf("target: %s\n", "mined a block")
		WriteABlock(genesis)
		break
	default:
		break
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
	target_fullAddr := fmt.Sprintf("%s", fullAddr)
	fmt.Printf("target: %s\n", target_fullAddr)
	time.Sleep(10 * time.Second)
	return target_fullAddr
}
