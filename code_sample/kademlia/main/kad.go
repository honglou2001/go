package main

import "fmt"

var node1 int=1  //0001
var node2 int=2  //0010
var node3 int=3  //0011
var node4 int=4  //0100

func main() {
	var i int ;
    i = 0x0001 ^ 0x0011

    //0001 ^ 0011 = 2;
    //0001 ^ 0010 = 3;
	//0001 ^ 0101 = 4

    //0011 ^ 0010 = 1
    //0011 ^ 0101 = 6

    //0011 ^ 0001 = 2


    //0010 ^ 0101 = 7

	fmt.Println(i);

}
