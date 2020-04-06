package main

import (
	"time"
)
import "yqx_go/code_sample/paxos"

func main() {
	done := make(chan int)

	go paxos.StartServer()
	time.Sleep(1000 * time.Millisecond)
	go paxos.RunClient()

	<-done
}
