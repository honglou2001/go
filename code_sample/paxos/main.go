package paxos

import (
	"time"
)

func main() {
	done := make(chan int)

	go startServer()
	time.Sleep(1000 * time.Millisecond)
	go runClient()

	<-done
}
