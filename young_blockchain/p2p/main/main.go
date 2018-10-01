package main

import (
	"fmt"
	"yqx_go/young_blockchain/p2p"
)

func main() {
	p2p.Run()
	fmt.Printf("TestStartRunner: %s\n", "ok")
}
