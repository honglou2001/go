package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestStartRunner(t *testing.T) {
	//first
	listentPort := 10001
	ha, fullAddr, err := initialAddress(listentPort)
	go func() {
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("StartRunner run \"go run main.go -l %d -d %s -secio\" on a different terminal\n", listentPort+1, fullAddr)
		StartRunner(ha, "")
	}()
	time.Sleep(5 * time.Second)
	//second
	listentPort2 := 10002
	ha2, fullAddr2, err2 := initialAddress(listentPort2)
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Printf("StartRunner2 run \"go run main.go -l %d -d %s -secio\" on a different terminal\n", listentPort2, fullAddr2)
	target2 := fmt.Sprintf("%s", fullAddr)
	fmt.Printf("target2: %s\n", target2)
	StartRunner(ha2,target2)

	fmt.Printf("TestStartRunner: %s\n", target2)
}
