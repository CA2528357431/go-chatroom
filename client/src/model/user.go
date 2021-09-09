package model

type Uesr struct {
	UesrId int `json:"uesrId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}


type UserRes struct {
	UserId int `json:"uesrId"`
	UserName string `json:"userName"`
}

type CLientUser struct {
	UserInfo UserRes
}