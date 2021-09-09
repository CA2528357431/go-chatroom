package model

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)


type UserDAO struct {
	Pool *redis.Pool
}

// 必须初始化
var (
	UserDao *UserDAO
)

// 初始化DAO
func Initud(pool *redis.Pool)  {
	UserDao = &UserDAO{Pool: pool}
}

func (this *UserDAO)GetUser(id int)(User,error)  {
	conn := this.Pool.Get()
	defer conn.Close()
	reply,err:=conn.Do("hget","users",id)
	if err!=nil{
		return User{}, err
	}
	// 由于没有用redis.string 因此必须手动判断获取是否为nil
	if reply==nil{
		return User{}, redis.ErrNil
	}

	data := reply.([]byte)
	var res User

	err = json.Unmarshal(data,&res)
	if err!=nil{
		return User{}, err
	}


	return res,nil

}

func (this *UserDAO)Login(id int,pwd string)(User,int)  {
	user,err := this.GetUser(id)
	var num int
	if err==redis.ErrNil {
		return User{}, 500
	}else if err!=nil{
		return User{}, 300
	}
	if pwd!=user.UserPwd{
		num = 501
	}
	return user,num
}

func (this *UserDAO)Register(UserId int,UserPwd string,UserName string )int  {

	user := User{
		UserId:   UserId,
		UserPwd:  UserPwd,
		UserName: UserName,
	}

	_,err := this.GetUser(UserId)
	if err==redis.ErrNil {

		data,_ := json.Marshal(user)

		str := string(data)

		conn := this.Pool.Get()
		conn.Do("hset","users",user.UserId,str)

		return 0

	}else if err!=nil{
		return 300
	}

	return 500
}



