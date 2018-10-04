package main

import "fmt"
import "yqx_go/young_blockchain/p2p"

var shouldQuit = make(chan struct{})

func main() {

	p2p.Run()
	fmt.Printf("TestStartRunner: %s\n", "running")

	for {
		select {
		case <-shouldQuit:
			//cleanUp()
			fmt.Printf("TestStartRunner: %s\n", "quiting")
			return
		 //default:
			//fmt.Printf("TestStartRunner: %s\n", time.Now())
		}
	}

	//再另外一个协程中，如果运行遇到非法操作或不可处理的错误，就向shouldQuit发送数据通知程序停止运行
	close(shouldQuit)
	fmt.Printf("TestStartRunner: %s\n", "quit all")
}
