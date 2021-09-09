package processor

import (
	"encoding/json"
	"fmt"
	"model"
	"utils"
)

var MyUsers = make(map[int] model.UserRes)
var CurUser model.CLientUser


func Set(data []byte)  {
	var notice utils.StatusNotify
	err := json.Unmarshal(data,&notice)
	if err!=nil{
		fmt.Println("Set error",err)
	}

	user := notice.User
	fmt.Println(user.UserId,user.UserName,notice.Status)


	MyUsers[user.UserId] = user


}

func ShowList(){
	for _,per := range MyUsers{
		fmt.Println(per.UserId,per.UserName)
	}
}