package processor

import (
	"fmt"
)

type UserManage struct {
	OnlineList map[int]*UserProcess
}

// 唯一的
var UserMg *UserManage

// 初始化用户管理器
func InitUserMg()  {
	UserMg = &UserManage{OnlineList: make(map[int]*UserProcess)}
}

func (this *UserManage)Set(one *UserProcess) {
	this.OnlineList[one.UserItem.UserId] = one
}

func (this *UserManage)Pop(id int) {
	delete(this.OnlineList,id)
}

func (this *UserManage)GetAll()map[int]*UserProcess {
	return this.OnlineList
}

func (this *UserManage)GetUser(id int)*UserProcess {
	res,err := this.OnlineList[id]
	if !err{
		fmt.Println("the User doesn't exist or offline")
	}
	return res
}

