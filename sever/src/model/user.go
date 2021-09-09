package model

type User struct {
	UserId int `json:"uesrId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type UserRes struct {
	UserId int `json:"uesrId"`
	UserName string `json:"userName"`
}


