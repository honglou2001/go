package main

import "fmt"

func consumer(ch chan string,done chan string)  {
	for i := range ch{
		fmt.Println("recive",i)
	}
	done<-"ok"

}
func product(ch chan string) {
	for i := 0; i<5; i++{
		//ch<-i
		s := fmt.Sprintf("writeData 2: %d,%s\n", -i,"uuu")
		ch<-s

	}
	close(ch)
}

func main() {
	fmt.Println("start")
	ch := make(chan string)
	done := make(chan string)
	//分别用goroutine启动两个协程：生产和消费
	go consumer(ch,done)
	go product(ch)
	//done是为了让主进程等待消费者结束
	<-done

}

