package p2p

import (
	"fmt"
	"log"
	"sync"
	"time"
)
//Run run test
func Run() {
	//first
	target2 := RunServer(10001,"")
	//second
	target3 := RunServer(10002,target2)
	//third
	RunServer(10003,target3)

	time.Sleep(60 * time.Second)
}
//RunServer run test
func RunServer(listentPort int,target string) (string){
	//first
	var server *p2pObject
	server = &p2pObject{listenFport:listentPort,mutex : &sync.Mutex{}}
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

//func RunClient2(listentPort int,target2 string) {
//	//second
//	var server2 *p2pObject
//	server2 = &p2pObject{listenFport:listentPort,mutex : &sync.Mutex{}}
//	ha2, fullAddr2, err2 := server2.InitialAddress()
//	go func() {
//		if err2 != nil {
//			log.Fatal(err2)
//		}
//		log.Printf("StartRunner2 run \"go run main.go -l %d -d %s -secio\" on a different terminal\n", listentPort, fullAddr2)
//		//target2 := fmt.Sprintf("%s", fullAddr)
//		fmt.Printf("target2: %s\n", target2)
//		server2.StartRunner(ha2, target2)
//	}()
//	time.Sleep(80 * time.Second)
//	fmt.Printf("TestStartRunner: %s\n", target2)
//}
