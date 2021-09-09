package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
}

// 收包
func (this *Transfer)ReadMsg() (Msg,error)  {
	lengthByte := make([]byte,8)
	n,err :=  this.Conn.Read(lengthByte)
	if err!=nil{
		fmt.Println("read length error",err)
		return Msg{},err
	}
	if n!=8{
		fmt.Println("receive length error","message lost")
		return Msg{}, err
	}

	lengthUint := binary.BigEndian.Uint64(lengthByte)
	length := int(lengthUint)
	// 解码出长度

	msgJsonByte := make([]byte,1024)
	n,err =  this.Conn.Read(msgJsonByte)
	if err!=nil{
		fmt.Println("receive msg error",err)
		return Msg{}, err
	}
	if length!=n{
		fmt.Println("receive msg error","message lost")
		return Msg{}, err
	}

	var message Msg
	err = json.Unmarshal(msgJsonByte[:n],&message)
	if err!=nil{
		fmt.Println("json error",err)
		return Msg{}, err
	}
	return message, nil

}


// 发包
func (this *Transfer)WriteMsg(data []byte)error{

	var length uint64 = uint64(len(data))
	lengthByte := make([]byte,8)
	binary.BigEndian.PutUint64(lengthByte,length)
	_,err := this.Conn.Write(lengthByte)
	if err!=nil{
		fmt.Println("write length error",err)
		return err
	}
	// 长度编码

	_,err = this.Conn.Write(data)
	if err!=nil{
		fmt.Println("write data error",err)
		return err
	}
	return nil

}
