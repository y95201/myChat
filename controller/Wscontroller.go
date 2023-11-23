/*
 * @Description:
 * @Author: Y95201
 * @Date: 2023-01-13 17:18:37
 * @LastEditors: y95201 957612196@qq.com
 * @LastEditTime: 2023-08-10 15:10:46
 */
package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myChat/models"
	_ "myChat/services/wsChat"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 变量定义初始化
var (
	wsUpgrader = websocket.Upgrader{}
	// TextMessage 表示文本消息类型
	TextMessage = websocket.TextMessage
	// BinaryMessage 表示二进制消息类型
	BinaryMessage = websocket.BinaryMessage
)

// 封装 WebSocket 消息发送的方法
func sendMessage(ws *websocket.Conn, mt int, message []byte) error {
	err := ws.WriteMessage(mt, message)
	if err != nil {
		log.Println("Error during message writing:", err)
	}
	return err
}
func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func WsHandle(gin *gin.Context) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upGrader.Upgrade(gin.Writer, gin.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		mt, raw_message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}

		// 开启关闭连接监听
		ws.SetCloseHandler(func(code int, text string) error {
			fmt.Println(code, text) // 断开连接时将打印code和text
			return nil
		})

		//处理逻辑
		data := string(raw_message)
		// Time := CurrentTime()
		types := gjson.Get(data, "type")
		if len(types.String()) != 0 {
			user_id := gjson.Get(data, "user_id").Int()
			//聊天处理
			if types.Int() == 1 {
				UserChatProcessing(data, int64(mt), user_id, ws)
			}
			//登陆处理
			if types.Int() == 2 {
				LoginBusinessProcessing(data, int64(mt), user_id, ws)
			}
			//刷新聊天记录
			if types.Int() == 3 {
				GetRecordStartConversation(data, int(mt), int(user_id), ws)
			}
			//推出用户让其他客服看到
			if types.Int() == 4 {
				customer := gjson.Get(data, "customer_id").Int()
				//查询最后一条数据
				id := models.GetTheLastPieceData(int(user_id))
				models.PressThePrimaryKeyModify(id, map[string]interface{}{
					"PId":   0,
					"PName": "",
					"state": 1,
				})
				//删除绑定
				models.ThinkChatQuickDelete(int(user_id))
				models.ThinkChatQuickDelete(int(customer))
			}
			//用户搜索
			if types.Int() == 5 {
				return
			}
			//用户群发 region,user_id,type,msg
			if types.Int() == 6 {
				return
			}
			//聊天分页
			if types.Int() == 7 {
				return
			}
			//未定义
			if types.Int() == 8 {
				return
			}
			//未定义
			if types.Int() == 9 {
				return
			}

			isUser := 0
			if user_id > 10 {
				isUser = 1
			}
			GetChatListOnlineservice(ws, mt, data, isUser)
		} else {
			message := []byte(`{"params":"PING11111"}`)
			err = sendMessage(ws, mt, message)
			if err != nil {
				return
			}
		}
	}
}

// 获取聊天列表及消息数
func GetChatListOnlineservice(ws *websocket.Conn, mt int, data string, isUser int) {
	//在查询在线客服用于发送数据
	OnlineService := *models.GetListOfOnlineCustomerService()
	customerID := gjson.Get(data, "customer_id").Int()
	if isUser == 1 {
		userID := gjson.Get(data, "user_id").Int()
		types := gjson.Get(data, "type").Int()
		// 用户登录推送在线客服
		if types != 1 {
			userListMessage, _ := json.Marshal(map[string]interface{}{"service": ServiceChatlistNews(userID)})
			sendMessage(ws, mt, userListMessage)
			if customerID == 0 {
				for _, value := range OnlineService {
					userListMessage, _ := json.Marshal(map[string]interface{}{
						"Userlist": UserChatlist(int64(value.Uid))})
					sendMessage(ws, value.Fid, userListMessage)
				}
			}
		}
	} else {
		userListMessage, _ := json.Marshal(map[string]interface{}{"Userlist": UserChatlist(gjson.Get(data, "user_id").Int())})

		if customerID != 0 {
			for _, value := range OnlineService {
				if value.Fid == mt {
					sendMessage(ws, value.Fid, userListMessage)
				} else {
					messageData := map[string]interface{}{"Userlist": UserChatlist(int64(value.Uid))}
					message, _ := json.Marshal(messageData)
					sendMessage(ws, TextMessage, message)
				}
			}
		} else {
			sendMessage(ws, mt, userListMessage)
		}
	}
}

func ServiceChatlistNews(userId int64) map[int][]map[string]interface{} {
	chatList := models.ObtainUserChatList(userId, 1)

	item := make(map[int][]map[string]interface{})

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
				item[v.PId] = append(item[v.PId], personMap)
			}
		}
	}
	content := make(map[int][]map[string]interface{})
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

type UserChatItem struct {
	UId       int    `json:"u_id"`
	UName     string `json:"u_name"`
	UAvatar   string `json:"u_avatar"`
	Company   string `json:"company"`
	Count     int    `json:"count"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func UserChatlist(userId int64) []UserChatItem {
	var chatLists = models.GetLatestMessagesByUserId(userId)
	var userChatItems []UserChatItem
	Content := ""
	Count := 0
	for _, chat := range chatLists {

		if strings.Contains(chat.Content, "<img") {
			Content = "[图片]"
		} else if strings.Contains(chat.Content, `{"id":`) {
			Content = "[订单]"
		} else {
			Content = chat.Content
		}

		Count = int(models.GetUserChatCount(int64(chat.UId), int64(chat.PId)))

		userChatItem := UserChatItem{
			UId:       chat.UId,
			UName:     chat.UName,
			UAvatar:   chat.Avatar,
			Company:   chat.Name,
			Count:     Count, // 在示例中暂时设置为0，具体值需要根据你的业务逻辑确定
			Content:   Content,
			CreatedAt: chat.CreatedAt,
		}
		userChatItems = append(userChatItems, userChatItem)
	}
	return userChatItems
}

func LoginBusinessProcessing(data string, mt int64, user_id int64, ws *websocket.Conn) {
	username := gjson.Get(data, "msg").String()
	userList := models.GetChatUserService(user_id)
	is_user := 1
	if user_id < 10 {
		is_user = 2
	}
	// 检查返回的切片是否为空
	if userList == nil || len(*userList) == 0 {
		// 用户不存在的处理逻辑
		data := &models.ThinkChatUser{
			Fid:    int(mt),
			Uid:    int(user_id),
			Uname:  username,
			IsUser: is_user,
		}
		models.ChatUserSaveService(data)
	} else {
		//用户在别处登陆后修改fd 然后让其下线
		models.ChatUserUpdatesService(user_id, map[string]interface{}{
			"Fid":    int(mt),
			"Uname":  username,
			"IsUser": is_user,
		})
		sendMessage(ws, (*userList)[0].Fid, []byte(`{"ReportError":"1","msg":"抢占登录"}`))
	}
}

func UserChatProcessing(data string, mt int64, user_id int64, ws *websocket.Conn) {
	msg := gjson.Get(data, "msg").String()
	active_user := models.GetUserByFieldValue("id,name,avatar,phone", int(user_id))
	if len(msg) != 0 {
		media := 1                       //默认文字信息
		is_base64 := func_is_base64(msg) //判断是否是图片
		if !is_base64 {
			msg = ProcessImage(msg) //是图片则处理保存上传
			media = 2               // 图片
		}
		state := 1
		msg = ubbReplace(msg) //表情替换

		receiveId := gjson.Get(data, "customer_id")
		if len(receiveId.String()) != 0 {
			PId := 0
			PName := ""
			if receiveId.Int() != 0 {
				PId = int(receiveId.Int())
				if PId < 9 {
					PName = fmt.Sprintf("%d号客服", PId)
				} else {
					Exists := models.ThinkChatQuickExists(int(user_id), PId)
					receiv := models.GetUserByFieldValue("id", PId)
					PName = receiv.Name
					opponent := models.GetChatUserFirst(int64(PId))
					mt := opponent.Fid
					if Exists && mt != 0 {
						state = 0 // 这里需要定义state变量
						sendMessage(ws, mt, []byte(msg))
					}
				}
			}

			data := &models.ThinkChat{
				UId:       int(active_user.ID),
				UName:     active_user.Name,
				PId:       PId,
				PName:     PName,
				Content:   msg,
				State:     state,
				CreatedAt: CurrentTime(),
				Media:     media,
				Url:       "",
			}
			models.ChatUserSaveMessageService(data)
		} else {
			data := &models.ThinkChat{
				UId:       int(active_user.ID),
				UName:     active_user.Name,
				PId:       0,
				PName:     "",
				Content:   msg,
				State:     state,
				CreatedAt: CurrentTime(),
				Media:     media,
				Url:       "",
			}
			models.ChatUserSaveMessageService(data)
		}
		sendMessage(ws, int(mt), []byte(msg))
	} else {
		message := []byte(`{"msg":"发送数据错误"}`)
		sendMessage(ws, int(mt), message)
	}
}

func func_is_base64(data string) bool {
	exe := strings.Replace(strings.Replace(strings.SplitN(data, ";", 2)[0], "/", ".", 1), "data:image", "", 1)
	return exe != ".jpeg" && exe != ".jpg" && exe != ".png"
}

// 替换表情
func ubbReplace(str string) string {
	re := regexp.MustCompile(`\[\wem:([0-9]*)\]`)
	str = re.ReplaceAllStringFunc(str, func(match string) string {
		// 提取 [wem:数字] 中的数字
		numStr := strings.TrimPrefix(match, "[wem:")
		numStr = strings.TrimSuffix(numStr, "]")

		// 构建替换后的 HTML 字符串
		return fmt.Sprintf("<img style=\"width:20px; height:20px\" src=\"https://chat.gangxinbao.cn/uploads/emoji/%s.png\" />", numStr)
	})

	return str
}

// 非图片返回内容否则写入图片并上传
func ProcessImage(data string) string {
	// 判断是否是 base64 数据
	isBase64 := strings.HasPrefix(data, "data:image")
	if !isBase64 {
		// 如果不是 base64 数据，则直接返回原数据
		return data
	}

	basePath := "/www/wwwroot/easy/public/"
	file := "uploads/chat/"
	fileName := fmt.Sprintf("%d", time.Now().UnixNano()) + ".png"
	filePath := filepath.Join(basePath+file, fileName)
	// 将base64字符串解码成字节切片
	imageData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return "" // 返回空字符串或者其他错误处理逻辑
	}

	// 将图片数据写入文件
	err = ioutil.WriteFile(filePath, imageData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return "" // 返回空字符串或者其他错误处理逻辑
	}

	imgName := filepath.Join(file, fileName)
	return fmt.Sprintf(`<img style="max-width:2rem;max-height:3rem;" class="upload-img" src="http://cdn.gangxinbao.cn/%s">`, imgName)
}

func GetRecordStartConversation(data string, mt int, user_id int, ws *websocket.Conn) {

	customerS := gjson.Get(data, "customer_id")
	customer := 0
	types := 2

	if len(customerS.String()) != 0 {
		customer = int(customerS.Int())
	}
	models.ThinkChatQuickDelete(int(user_id))

	models.ThinkChatQuickSave(&models.ThinkChatBinding{
		Uid:  int(user_id),
		Ufd:  int(mt),
		Pid:  customer,
		Type: types,
	})
	//获取聊天记录
	var chatList []models.ThinkChat
	if user_id > 10 && customer < 10 { //用户和客服聊天
		chatList = models.GetUserAndServiceChatList(int(user_id))
	}
	if user_id < 10 && customer > 10 { //客服和用户聊天

		chatList = models.GetServiceAndUserChatList(int(customer))
	}
	if user_id > 10 && customer > 10 { //用户和用户聊天
		chatList = models.GetUserAndUserChatList(int(user_id), int(customer))
	}
	// 判断 chatList 是否为空
	if len(chatList) != 0 {
		var im string
		var timeString string
		for _, value := range chatList {
			dateTime, _ := time.Parse("2006-01-02 15:04:05", value.CreatedAt)
			if time.Now().Before(dateTime) {
				if im == "" {
					im = dateTime.Format("15")
					timeString = dateTime.Format("15:04:05")
				} else {
					if im != dateTime.Format("15") {
						im = dateTime.Format("15")
						timeString = dateTime.Format("15:04:05")
					}
				}
			} else {
				if im == "" {
					im = dateTime.Format("02")
					timeString = dateTime.Format("2006-01-02")
				} else {
					if im != dateTime.Format("02") {
						im = dateTime.Format("02")
						timeString = dateTime.Format("2006-01-02")
					}
				}
			}
			// fmt.Println(timeString)
			if len(timeString) != 0 {
				sendMessage(ws, int(mt), []byte(fmt.Sprintf(`{"system":"1","msg":"%s"}`, timeString)))
			}

			//遍历发送数据库历史消息
			SendUser := models.GetUserByFieldValue("id,name,avatar,phone", int(value.UId))
			AceptUsers := models.GetUserByFieldValue("id,name,avatar,phone", int(value.PId))
			messageJSON, _ := json.Marshal(map[string]interface{}{
				"p_id":       AceptUsers.ID,
				"p_name":     AceptUsers.Name,
				"p_avatar":   AceptUsers.Avatar,
				"u_id":       SendUser.ID,
				"userName":   SendUser.Name,
				"userAvatar": SendUser.Avatar,
				"created_at": value.CreatedAt,
				"avatar":     SendUser.Avatar,
				"msg":        value.Content,
				"media":      value.Media,
			})
			sendMessage(ws, int(mt), messageJSON)

			//修改消息已读
			if int(AceptUsers.ID) == int(user_id) {
				if value.State == 1 {
					models.PressThePrimaryKeyModify(value.ID, map[string]interface{}{
						"state": 0,
					})
				}
			}
			//若是客服则绑定消息对话为客服
			if value.PId == 0 {
				username := ""
				if user_id < 10 {
					username = fmt.Sprintf("%d号客服", user_id)
				}
				models.PressThePrimaryKeyModify(value.ID, map[string]interface{}{
					"p_id":   0,
					"p_name": username,
				})
			}
		}

	} else {
		if customer < 10 {
			user := models.GetUserByFirstValue("name", int(user_id))
			content := "您好" + user.Name + "，我是钢信宝平台客服，请简述您需要咨询的问题，我将第一时间为您解答！（全国服务热线：400-699-0208）[^_^]"
			sendMessage(ws, int(mt), []byte(content))
		} else {
			//无聊天内容时 发送欢迎语
			UserCompanyInfo := models.GetUserCompanyRelatedInquiry(int(user_id))
			if UserCompanyInfo.CompanyName != "" {
				sendMessage(ws, int(mt), []byte(fmt.Sprintf(`{"system":"1","msg":"%s"}`, CurrentTime())))
				content := fmt.Sprintf(`尊敬的客户您好！我是%s的业务员%s，主要负责：%s，您可以联系我的电话：%d`, UserCompanyInfo.CompanyName, UserCompanyInfo.Title, UserCompanyInfo.MainRespFor, UserCompanyInfo.Phone)
				message := fmt.Sprintf(`{"p_id":%d,"p_name":"%s","avatar":"%s","msg":"%s","created_at":"%s","media":1}`, customer, UserCompanyInfo.Title, UserCompanyInfo.Avatar, content, CurrentTime())
				sendMessage(ws, int(mt), []byte(message))
			}
		}
	}
}
