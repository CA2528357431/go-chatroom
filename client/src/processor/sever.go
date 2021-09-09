package processor

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"utils"
)

func ShowMenu() {

	for {
		fmt.Println("-------1. show online list---------")
		fmt.Println("-------2. send message---------")
		fmt.Println("-------3. message list---------")
		fmt.Println("-------4. exit system---------")
		var key int

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("online lise")
			ShowList()
		case 2:
			fmt.Println("please input what you want to say")
			var inputReader *bufio.Reader
			inputReader = bufio.NewReader(os.Stdin)
			str,_ := inputReader.ReadString('\n')
			content :=  strings.Trim(str,"\n")

			sp := SmsProcess{}
			sp.SendToAll(content)
		case 3:
			fmt.Println("msg list")
		case 4:
			fmt.Println("exit")
			os.Exit(0)
		default:
			fmt.Println("error input")
		}
	}

}

func ContactSever(conn net.Conn)  {
	tf := utils.Transfer{Conn: conn}
	for{

		msg,err := tf.ReadMsg()
		if err != nil{
			fmt.Println("sever error",err)
			return
		}

		switch msg.MsgType {
		case utils.StatusNotifyType:
			Set(msg.Data)
		case utils.SmsToAllType:
			sp := SmsProcess{}
			sp.Receive(msg.Data)
		}

	}

}


