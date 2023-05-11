/*
 * @Description:
 * @Author: Y95201
 * @Date: 2023-01-13 17:18:37
 * @LastEditors: Y95201
 * @LastEditTime: 2023-04-14 18:03:16
 */
package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"log"
	"myChat/models"
	_ "myChat/services/wsChat"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 变量定义初始化
var (
	wsUpgrader = websocket.Upgrader{}
	offline    = make(chan *websocket.Conn)
)

func WsHandle(gin *gin.Context) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upGrader.Upgrade(gin.Writer, gin.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		//拿到发送来的数据
		data := string(message)
		fmt.Println("clients message " + data)
		types := gjson.Get(data, "type")
		if len(types.String()) != 0 {
			//user_id := gjson.Get(data, "user_id")
			isUser := 1
			//if user_id.Int() < 10 {
			//	isuser = 0
			//}
			GetChatListOnlineservice(data, isUser)
			//m, ok := gjson.Parse(data).Value().(map[string]interface{})
			//if !ok {
			//	// not a map
			//}
			//fmt.Println(m)
		} else {
			message = []byte(`{"params":"PING11111"}`)
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
	//wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	//ws, err := upGrader.Upgrade(gin.Writer, gin.Request, nil)
	//defer ws.Close()
	//go read(ws)
	//select {}
}

//func read(c *websocket.Conn) {
//	for {
//		//读取ws中的数据
//		mt, message, err := c.ReadMessage()
//		if err != nil {
//			offline <- c
//			log.Println("ReadMessage error1", err)
//			return
//		}
//		//拿到发送来的数据
//		data := string(message)
//		fmt.Println("clients message " + data)
//		types := gjson.Get(data, "type")
//		if len(types.String()) != 0 {
//			//user_id := gjson.Get(data, "user_id")
//			isUser := 1
//			//if user_id.Int() < 10 {
//			//	isuser = 0
//			//}
//			GetChatListOnlineservice(data, isUser)
//
//		} else {
//			message = []byte(`{"params":"PING11111"}`)
//		}
//		err = c.WriteMessage(mt, message)
//		if err != nil {
//			break
//		}
//	}
//}

type isUser int

// 获取聊天列表及消息数
func GetChatListOnlineservice(data string, isUser int) {
	OnlineUser := models.GetListOfOnlineUsers()
	//OnlineService := models.GetListOfOnlineCustomerService()
	if isUser == 1 {
		user_id := gjson.Get(data, "user_id").Int()
		//fmt.Println(reflect.TypeOf(user_id))
		ServiceChatlistNews(user_id)
		//log.Println(OnlineService)
	} else {
		log.Println(OnlineUser)
	}
}

type ChatItem struct {
	Uid       int    `json:"u_id"`
	Uname     string `json:"u_name"`
	Pid       int    `json:"p_id"`
	Pname     string `json:"p_name"`
	Content   string `json:"content"`
	State     int    `json:"state"`
	Media     int    `json:"media"`
	CreatedAt string `json:"created_at"`
	Avatar    string `json:"avatar"`
}

func ServiceChatlistNews(userId int64) {
	chatList := models.ObtainUserChatList(userId)
	//item := make(map[int][]ChatItem)
	for _, v := range chatList {
		log.Println(v)
		//if _, ok := item[v.Uid]; !ok {
		//} else {
		//	item[v.Uid] = append(item[v.Uid], v)
		//}
	}
}
