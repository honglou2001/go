package main

import (
	"bufio"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type BlockTest struct {
	Height int
}
func produce(p chan<- BlockTest) {

	b := &BlockTest{1}
	p <- *b
	//for i := 0; i < 10; i++ {
	//	p <- i
	//	fmt.Println("pro1:", i)
	//}
}

//func produce2(p chan<- BlockTest) {
//	for i := 10; i < 20; i++ {
//		p <- i
//		fmt.Println("pro2:", i)
//	}
//}

func consumer1(c <-chan BlockTest) {
	for {
		v := <-c
		spew.Dump(v)
		fmt.Println("receive:", v.Height,time.Now())
	}
}


//func consumer1(c <-chan BlockTest) {
//	for {
//		v := <-c
//		fmt.Println("receive:", v,time.Now())
//	}
//}
var Ch = make(chan BlockTest, 3)
func main() {
	//ch := make(chan int)
	go produce(Ch)
	go consumer1(Ch)
	time.Sleep(10 * time.Second)
	//go produce2(Ch)
	//time.Sleep(10 * time.Second)
	//Ch <- 201
	//time.Sleep(10 * time.Second)
	//Ch <- 202
	//time.Sleep(10 * time.Second)
	//Ch <- 203
	//ch <- 202
	fmt.Println("receive:", "endend")
	fmt.Println("Please input height of a block: ")
	stdReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		sendData = strings.Replace(sendData, "\n", "", -1)
		bpm, err := strconv.Atoi(sendData)

		if(bpm>1){
			b := &BlockTest{bpm}
			Ch <- *b
			//Ch <- bpm
		}
	}

	//n := flag.Int("n", 1, "enable secio")
	//flag.Parse()
	//if *n > 1 {
	//	log.Fatal("Please provide a port to bind on with -l")
	//}
}

