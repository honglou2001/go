package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "test.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}
	defer file.Close() //关闭文件
	/*	reader := bufio.NewReader(file)    //带缓冲区的读写
		for {
			str, err := reader.ReadString('\n')    // 循环读取一行
			if err != nil {
				fmt.Println("read string failed, err: ", err)
				return
			}
			fmt.Println("read string is %s: ", str)
		}*/

	err_remove := os.Remove(fileName) //删除文件test.txt
	if err_remove != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("file remove Error!")
		//输出错误详细信息
		fmt.Printf("%s", err_remove)
	} else {
		//如果删除成功则输出 file remove OK!
		fmt.Print("file remove OK!")
	}
}
