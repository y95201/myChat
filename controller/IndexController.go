/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-12-20 14:00:12
 * @LastEditors: Y95201
 * @LastEditTime: 2022-12-20 14:43:48
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"log"
	"myChat/models"
	"myChat/services/user_service"
	"net/http"
	"strings"
)

func Index(c *gin.Context) {
	// 已登录跳转界面
	//userInfo := user_service.GetUserInfo(c)
	//if len(userInfo) > 0 {
	//	c.Redirect(http.StatusFound, "/home")
	//	return
	//}

	//OnlineUserCount := primary.OnlineUserCount()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "聊天室",
	})
}
func Login(c *gin.Context) {
	user_service.Login(c)
}

func Logout(c *gin.Context) {
	user_service.Logout(c)
}
func CeShi(c *gin.Context) {
	mgs := `{"type":2,"user_id":"4","msg":"9号客服","customer_id":""}`
	types := gjson.Get(mgs, "type")
	if len(types.String()) != 0 {
		isUser := 0
		content := GetChatListOnlineserviceS(mgs, isUser)
	} else {

	}
	//content := OnlineService(4, 1)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  content,
	})
	return
}
func GetChatListOnlineserviceS(data string, isUser int) {
	OnlineUser := models.GetListOfOnlineUsers()
	OnlineService := models.GetListOfOnlineCustomerService()
	if isUser == 1 {
		//用户登录推送在线客服

		log.Println(OnlineService)
		log.Println(OnlineUser)
	} else {
		log.Println(gjson.Get(data, "customer_id").Int())
		if len(gjson.Get(data, "customer_id").String()) == 0 {

		} else {

		}
	}
}
func OnlineService(userId int64, isUser int) {
	//$onlineUser = $this->toArrays(ThinkChatUserModel::create()->where('u_id','10','>')->all());
	//$onlineService = $this->toArrays(ThinkChatUserModel::create()->where('u_id','10','<')->all());
	if isUser == 1 {

	} else {

	}
	UserChatlist(userId)
	ServiceChatlistNewsY(userId)
}

func UserChatlist(userId int64) map[int][]map[string]interface{} {
	var chatList = models.ObtainUserChatList(int64(userId), 2)
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
			item[v.UId] = []map[string]interface{}{personMap}
		} else {
			item[v.UId] = append(item[v.UId], personMap)
		}
	}
	return item
}

func ServiceChatlistNewsY(userId int64) map[int][]map[string]interface{} {
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
	for key, _ := range item {

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
