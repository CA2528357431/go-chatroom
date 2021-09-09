package processor

import (
	"encoding/json"
	"fmt"
	"model"
	"net"

)
import "utils"

type UserProcess struct {
	Conn net.Conn
	UserItem model.User
}

func (this *UserProcess)Register(data []byte){
	var registerMes utils.RegisterMes
	err := json.Unmarshal(data,&registerMes)
	if err!=nil{
		fmt.Println("Register error",err)
		return
	}
	// 解包

	neo := utils.RegisterRes{}
	num := model.UserDao.Register(registerMes.UserId,registerMes.UserPwd,registerMes.UserName)
	neo.Code = num

	if num==0{
		fmt.Println(registerMes.UserId,"register successfully")
	}else if num==500{
		neo.Err = utils.REGISTER_WRONG_ID
	}else if neo.Code == 300{
		neo.Err=utils.REGISTER_NETWORK_ERROR
	}


	neoData,err := json.Marshal(neo)
	if err!=nil{
		fmt.Println("Login error",err)
		return
	}


	res := utils.Msg{
		MsgType: utils.RegisterResType,
		Data:    neoData,
	}

	// 写包

	resByte,err := json.Marshal(res)

	tf := utils.Transfer{Conn: this.Conn}

	err = tf.WriteMsg(resByte)
	if err!=nil{
		fmt.Println("Login error",err)
		return
	}

}

func (this *UserProcess)Login(data []byte)  {

	var loginStruct utils.LoginMes
	err := json.Unmarshal(data,&loginStruct)
	if err!=nil{
		fmt.Println("Login error",err)
		return
	}

	neo := utils.LoginRes{}
	user,num := model.UserDao.Login(loginStruct.UserId,loginStruct.UserPwd)
	neo.Code = num

	if num==0{
		// 为manage提供信息
		this.UserItem = user
		this.Notify(utils.ONLINE)


		users := make([]model.UserRes,0,16)
		for key,value := range UserMg.OnlineList{
			per := model.UserRes{
				UserId:   key,
				UserName: value.UserItem.UserName,
			}
			users = append(users,per)
		}
		neo.Users = users
		neo.User = model.UserRes{
			UserId:   user.UserId,
			UserName: user.UserName,
		}
		UserMg.Set(this)

		fmt.Println(loginStruct.UserId,"online")


	}else if num==500{
		neo.Err = utils.LOGIN_WRONG_ID
	}else if neo.Code == 501 {
		neo.Err=utils.LOGIN_WRONG_PWD
	}else if neo.Code == 300{
		neo.Err=utils.LOGIN_NETWORK_ERROR
	}


	neoData,err := json.Marshal(neo)
	if err!=nil{
		fmt.Println("Login error",err)
		return
	}


	res := utils.Msg{
		MsgType: utils.LoginResType,
		Data:    neoData,
	}

	resByte,err := json.Marshal(res)
	if err!=nil{
		fmt.Println("Login error",err)
		return
	}

	tf := utils.Transfer{Conn: this.Conn}

	err = tf.WriteMsg(resByte)
	if err!=nil{
		fmt.Println("Login error",err)
		return
	}
}

func (this *UserProcess)NotifyToPer(user model.UserRes,status string)  {
	neo := utils.StatusNotify{
		User:   user,
		Status: status,
	}
	neoData,err := json.Marshal(neo)
	if err!=nil{
		fmt.Println("Notify error",err)
		return
	}

	res := utils.Msg{
		MsgType: utils.StatusNotifyType,
		Data:    neoData,
	}

	resByte,err := json.Marshal(res)
	if err!=nil{
		fmt.Println("Notify error",err)
		return
	}

	tf := utils.Transfer{Conn: this.Conn}

	err = tf.WriteMsg(resByte)
	if err!=nil{
		fmt.Println("Notify error",err)
		return
	}

}

func (this *UserProcess)Notify(status string) {
	for _,p := range UserMg.OnlineList{
		userRes := model.UserRes{
			UserId:   this.UserItem.UserId,
			UserName: this.UserItem.UserName,
		}
		p.NotifyToPer(userRes,status)
	}
}