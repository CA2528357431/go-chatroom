package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"utils"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendToAll(content string)  {

	sms := utils.SmsToAll{
		Content: content,
		From:    CurUser.UserInfo,
	}

	data,err := json.Marshal(sms)

	if err!=nil{
		fmt.Println("SendToAll error",err)
		return
	}

	msg := utils.Msg{
		MsgType: utils.SmsToAllType,
		Data:    data,
	}

	conn,err := net.Dial("tcp","0.0.0.0:8888")
	if err!=nil{
		fmt.Println("login error",err)
		return
	}
	defer conn.Close()

	tf := utils.Transfer{Conn: conn}

	msgByte,err := json.Marshal(msg)
	if err!=nil{
		fmt.Println("SendToAll error",err)
		return
	}

	err = tf.WriteMsg(msgByte)
	if err != nil {
		fmt.Println("SendToAll error",err)
		return
	}


}

func (this *SmsProcess) Receive(data []byte)  {
	var sms utils.SmsToAll
	err := json.Unmarshal(data,&sms)
	if err!=nil{
		fmt.Println("Receive error",err)
	}

	fmt.Println(sms.From.UserId,sms.From.UserName,sms.Content)

}