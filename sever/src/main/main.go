package main

import (
	"fmt"
	"model"
	"net"
	"processor"
	"time"

)



// 初始化
func init()  {
	Initpool("0.0.0.0:6379",16,time.Second*60)
	model.Initud(pool)
	processor.InitUserMg()
}

func main() {
	fmt.Println("--------------多人聊天系统后台------------")
	listener,err := net.Listen("tcp","0.0.0.0:8888")
	if err!=nil{
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for {
		conn,err := listener.Accept()
		if err!=nil{
			fmt.Println("listen error",err)
			return
		}


		go process(conn)
	}

}

// 调用主控
func process(conn net.Conn)  {
	mp := MainProcess{Conn: conn}
	err := mp.mainProcess()
	if err!=nil{
		fmt.Println("process error",err)
		return
	}
}