package main
import (
	"fmt"
	"os"
	"processor"
)

//定义两个变量，一个表示用户id, 一个表示用户密码
var userId int
var userPwd string
var userName string

func main() {

	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	//var loop = true

	for {
		fmt.Println("----------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1 :
			fmt.Println("登陆聊天室")
			up := processor.UserProcess{}
			up.Login()
			break
		case 2 :
			fmt.Println("注册用户")
			up := processor.UserProcess{}
			up.Register()
		case 3 :
			fmt.Println("退出系统")
			os.Exit(0)
		default :
			fmt.Println("你的输入有误，请重新输入")
		}

	}

}