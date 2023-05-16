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
	"strings"
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

func ServiceChatlistNews(userId int64) map[int][]map[string]interface{} {
	chatList := models.ObtainUserChatList(userId)
	item := make(map[int][]map[string]interface{})
	content := make(map[int][]map[string]interface{})
	for _, v := range chatList {
		personMap := map[string]interface{}{
			"Uid":       v.UId,
			"Uname":     v.UName,
			"Pid":       v.PId,
			"Pname":     v.PName,
			"Content":   v.Content,
			"State":     v.State,
			"Media":     v.Media,
			"CreatedAt": v.CreatedAt,
			"Avatar":    v.Avatar,
			"Name":      v.Name,
		}
		if _, ok := item[v.UId]; !ok {
			if v.UId < 10 {
				personMap["Uname"] = "钢信宝客服"
				personMap["Avatar"] = "images/IN7gUqUPXXK2AGgepnGVk1fq5rVRZj7NqCSXO4NB.png"
				item[0] = append(item[0], personMap)
			} else {
				item[v.UId] = []map[string]interface{}{personMap}
			}
		} else {
			if v.PId > 10 && v.PId != int(userId) {
				item[v.UId] = append(item[v.UId], personMap)
			}
		}
	}
	for key := range item {
		if key != int(userId) {
			data := models.ChatLastPieceData(int64(userId), key, 2)
			if key < 10 {
				data = models.ChatLastPieceData(int64(userId), key, 1)
			}

			value := map[string]interface{}{}
			if strings.Contains(data.Content, "<img") {
				value["Content"] = "[图片]"
			} else if strings.Contains(data.Content, `{"id":`) {
				value["Content"] = "[订单]"
			} else {
				value["Content"] = data.Content
			}
			value["Count"] = 0
			if int(userId) == data.PId {
				value["Count"] = models.CountUserMessages(int64(userId), key)
			}

			value["UID"] = key
			if key < 10 {
				value["UAvatar"] = "images/IN7gUqUPXXK2AGgepnGVk1fq5rVRZj7NqCSXO4NB.png"
				value["UName"] = "客服"
				value["Phone"] = ""
			} else {
				user := models.GetUserByFirstValue("name,avatar,phone", key)
				value["UAvatar"] = user.Avatar
				value["UName"] = user.Name
				value["Phone"] = user.Phone
			}
			value["State"] = data.State
			value["created_at"] = data.CreatedAt
			value["media"] = data.Media
			content[key] = append(content[key], value)
		}
	}
	return content
}
