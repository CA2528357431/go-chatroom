package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"utils"
)

type UserProcess struct {
}

func (this *UserProcess)Register()  {
	var id int
	fmt.Print("ID: ")
	fmt.Scanf("%d\n",&id)

	var pwd string
	fmt.Print("PWD: ")
	fmt.Scanf("%s\n",&pwd)

	var name string
	fmt.Print("NAME: ")
	fmt.Scanf("%s\n",&name)

	conn,err := net.Dial("tcp","0.0.0.0:8888")
	if err!=nil{
		fmt.Println("register error",err)
		return
	}
	defer conn.Close()

	registerMsg := utils.RegisterMes{
		UserId:   id,
		UserPwd:  pwd,
		UserName: name,
	}

	data,err := json.Marshal(registerMsg)
	if err!=nil{
		fmt.Println("register error",err)
		return
	}

	msg := utils.Msg{
		MsgType: utils.RegisterMesType,
		Data:    data,
	}

	msgByte,err := json.Marshal(msg)
	if err!=nil{
		fmt.Println("register error",err)
		return
	}

	tf := utils.Transfer{Conn: conn}

	err = tf.WriteMsg(msgByte)
	if err!=nil{
		fmt.Println("register error",err)
		return
	}

	neomsg,err := tf.ReadMsg()
	if err!=nil{
		fmt.Println("register error",err)
		return
	}

	neodata := neomsg.Data
	var res utils.RegisterRes
	err = json.Unmarshal(neodata, &res)

	if err!=nil{
		fmt.Println("register error",err)
		return
	}

	if res.Code==0{
		fmt.Println("register successfully")
	}else {
		fmt.Println(res.Code,res.Err)
	}

}


func (this *UserProcess) Login()  {
	var id int
	fmt.Print("ID: ")
	fmt.Scanf("%d\n",&id)

	var pwd string
	fmt.Print("PWD: ")
	fmt.Scanf("%s\n",&pwd)

	conn,err := net.Dial("tcp","0.0.0.0:8888")
	if err!=nil{
		fmt.Println("login error",err)
		return
	}

	loginMsg := utils.LoginMes{
		UserId:  id,
		UserPwd: pwd,
	}

	data,err := json.Marshal(loginMsg)
	if err!=nil{
		fmt.Println("login error",err)
		return
	}

	msg := utils.Msg{
		MsgType: utils.LoginMesType,
		Data:    data,
	}

	msgByte,err := json.Marshal(msg)
	if err!=nil{
		fmt.Println("login error",err)
		return
	}

	tf := utils.Transfer{Conn: conn}
	err = tf.WriteMsg(msgByte)
	if err!=nil{
		fmt.Println("login error",err)
		return
	}



	neomsg,err := tf.ReadMsg()
	if err!=nil{
		fmt.Println("login error",err)
		return
	}

	neodata := neomsg.Data
	var res utils.LoginRes
	err = json.Unmarshal(neodata, &res)

	if err!=nil{
		fmt.Println("login error",err)
		return
	}

	if res.Code==0{
		fmt.Printf("login successfully, %s\n",res.User.UserName)
		fmt.Println("online list")
		for _,per := range res.Users{
			fmt.Println(per.UserId,":",per.UserName)
			MyUsers[per.UserId] = per
		}

		// 初始化CurUser
		CurUser.UserInfo = res.User


		go ContactSever(conn)
		ShowMenu()

	}else {
		fmt.Println(res.Code,res.Err)
	}


}




