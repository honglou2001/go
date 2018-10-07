package main

import (
	"fmt"
)

// 实现枚举例子

type State int

// iota 初始化后会自动递增
const (
	Running State = iota // value --> 0
	Stopped              // value --> 1
	Rebooting            // value --> 2
	Terminated           // value --> 3
)

func (this State) String() string {
	switch this {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	default:
		return "Unknow"
	}
}

func main() {
	state12 := Stopped
	fmt.Println("state", state12)

	s1 :=Stopped.String()
	fmt.Println("s1", s1)
}
// 输出 state Running
// 没有重载String函数的情况下则输出 state 0
