package paxos

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
	"strings"
)

//服务端输入参数
type Args struct {
	StrParams string
}

//服务端输出参数
type Reply struct {
	StrResult string
}

//服务的结构体
type PaxosService struct {
	//进程参数
	Host string `本进程所在主机名或者ip地址`
	Port string `本进程使用的端口号`

	//paxos参数
	T_max   int    `当前已经发布的最大票号`
	C       string `当前存储的命令`
	T_store int    `存储命令C的票`
}

//服务的测试函数
func (ps *PaxosService) Test1(args *Args, reply *Reply) error {
	reply.StrResult = args.StrParams + "123"
	return nil
}

func (ps *PaxosService) Process(args *Args, reply *Reply) error {
	reply.StrResult = args.StrParams + "123"
	//切分参数
	params := strings.Split(args.StrParams, ":")

	//输出一下
	//	for _, p := range params {
	//		fmt.Println(p)
	//	}
	//根据参数情况进行回复
	switch params[2] {
	case "require_ticket":
		//阶段1 客户端请求票
		fmt.Println("阶段1 客户端请求票")
		require_t, err := strconv.Atoi(params[3])
		if err != nil {
			fmt.Println("字符串转整型失败", err, params[3])
			return err
		}
		fmt.Println("阶段1 ,t=", require_t)
		if require_t > ps.T_max {
			ps.T_max = require_t
			reply.StrResult = ps.Host + ":" + ps.Port + ":" + "response_ticket" + ":" + "ok" + ":" + ps.C
		}
	case "require_propose":
		fmt.Println("阶段2 客户端请求propose")
		require_t, err := strconv.Atoi(params[3])
		if err != nil {
			fmt.Println("字符串转整型失败", err, params[3])
			return err
		}
		fmt.Println("require_t=", require_t, ", ps.T_max=", ps.T_max)
		if require_t == ps.T_max {
			ps.C = params[4]
			ps.T_store = require_t
			reply.StrResult = ps.Host + ":" + ps.Port + ":" + "response_propose" + ":" + "success"
		}
	case "require_commit":
		fmt.Println("阶段3 客户端请求commit")
		require_t, err := strconv.Atoi(params[3])
		if err != nil {
			fmt.Println("字符串转整型失败", err, params[3])
			return err
		}
		if require_t == ps.T_store {
			if params[4] == ps.C {
				reply.StrResult = ps.Host + ":" + ps.Port + ":" + "response_propose" + ":" + "run" + ":" + params[4]
			}
		}
	default:
		fmt.Println("命令不明，无法确认阶段，返回fail")
		reply.StrResult = ps.Host + ":" + ps.Port + ":" + "fail"
	}

	return nil
}

//启动服务
func startServer() {
	fmt.Println("start server...")
	ps := PaxosService{
		Host:    "127.0.0.1",
		Port:    "1234",
		T_max:   0,
		C:       "none",
		T_store: 0,
	}
	server := rpc.NewServer()
	server.Register(&ps)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("server running...")
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}
