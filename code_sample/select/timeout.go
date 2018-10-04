package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:  //如果有数据，下面打印。但是有可能ch一直没数据
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second): //上面的ch如果一直没数据会阻塞，那么select也会检测其他case条件，检测到后3秒超时
				fmt.Println("超时")
				quit <- true  //写入
			}
		}

	}() //别忘了()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit //这里暂时阻塞，直到可读
	fmt.Println("程序结束")

}

//防止读取超时
//select {
//case <- time.After(time.Second*2):
//println("read channel timeout")
//case i := <- ch:
//println(i)
//}
//
////防止写入超时
//select {
//case <- time.After(time.Second *2):
//println("write channel timeout")
//case ch <- "hello":
//println("write ok")
//}