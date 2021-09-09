package processor

import (
	"encoding/json"
	"fmt"
	"utils"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendToAll(data []byte)  {


	msg := utils.Msg{
		MsgType: utils.SmsToAllType,
		Data:    data,
	}

	for _,p := range UserMg.OnlineList{

		tf := utils.Transfer{Conn: p.Conn}
		msgByte,err := json.Marshal(msg)
		if err!=nil{
			fmt.Println("SendToAll error",err)
		}
		tf.WriteMsg(msgByte)
	}

}
