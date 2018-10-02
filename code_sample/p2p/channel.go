package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func consumer(ch chan string,done chan string) {

	for{
		var j= <-ch
		fmt.Println("recive", j)
	}

	fmt.Println("done", "d")

	//for i := range ch{
	//	fmt.Println("recive",i)
	//}
	//done<-"ok"

}
func product(ch chan string) {
	for i := 0; i<5; i++{
		time.Sleep(2 * time.Second)
		s := fmt.Sprintf("writeData 2: %d,%s\n", -i,"uuu")
		ch<-s

	}
	wg.Add(1)
	wg.Wait()
	//close(ch)
}

func product2(ch chan string) {
	for i := 0; i<5; i++{
		time.Sleep(1 * time.Second)
		s := fmt.Sprintf("writeData 2: %d,%s\n", -i,"222")
		ch<-s

	}
	//close(ch)
}

func main() {

	wg.Add(1)
	fmt.Println("start")
	ch := make(chan string,3)
	done := make(chan string)
	//分别用goroutine启动两个协程：生产和消费
	go consumer(ch,done)
	go product(ch)

	time.Sleep(30 * time.Second)
	fmt.Println("recive", "ssss")
	//go product2(ch)
	//done是为了让主进程等待消费者结束
	//
	//for i := 0; i<5; i++{
	//	time.Sleep(1 * time.Second)
	//	s := fmt.Sprintf("main : %d,%s\n", -i,"kkkk")
	//	ch<-s
	//
	//}

	wg.Wait()
	fmt.Println("recive", "endend")
	//<-done

}

