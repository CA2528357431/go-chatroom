package utils

import "model"

const (
	LoginMesType = "LoginMes"
	LoginResType = "LoginRes"
	RegisterMesType = "RegisterMes"
	RegisterResType = "RegisterRes"
	StatusNotifyType = "StatusNotify"
	SmsToAllType = "SmsToAll"
	SmsToPerType = "SmsToPer"
)



type Msg struct {
	MsgType string `json:"msgType"`
	Data []byte `json:"data"`
}

type LoginMes struct{
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
}

type LoginRes struct {
	Code int `json:"code"`
	// 0代表没问题
	Err string `json:"err"`
	User model.UserRes `json:"user"`
	Users []model.UserRes `json:"users"`
}

type RegisterMes struct{
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"user_name"`
}

type RegisterRes struct {
	Code int `json:"code"`
	// 0代表没问题
	Err string `json:"err"`
}

const (
	ONLINE = "online"
	OFFLINE = "offline"
)

type StatusNotify struct {
	User model.UserRes
	Status string
}

type SmsToPer struct {
	Content string
	Target model.UserRes
	From model.UserRes
}

type SmsToAll struct {
	Content string
	From model.UserRes
}