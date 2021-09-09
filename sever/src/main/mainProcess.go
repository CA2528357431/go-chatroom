package main

import (
	"fmt"
	"net"
	"processor"
	"utils"
)

type MainProcess struct {
	Conn net.Conn
}


func (this *MainProcess)mainProcess()error  {

	tf := utils.Transfer{Conn: this.Conn}

	message,err:=tf.ReadMsg()
	if err!=nil{
		fmt.Println("mainProcess error",err)
		return err
	}

	switch message.MsgType {
	case utils.LoginMesType:
		up := processor.UserProcess{Conn: this.Conn}
		up.Login(message.Data)
	case utils.RegisterMesType:
		up := processor.UserProcess{Conn: this.Conn}
		up.Register(message.Data)
	case utils.SmsToAllType:

		sp := processor.SmsProcess{}
		sp.SendToAll(message.Data)
	default:
		fmt.Println(message.MsgType)
		
	}

	return nil
}