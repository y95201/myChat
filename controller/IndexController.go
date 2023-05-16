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
	"myChat/models"
	"myChat/services/user_service"
	"net/http"
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
	content := ServiceChatlistNewsY(4)
	//Data := arraySort(content, 'CreatedAt')
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  content,
	})
	return
}
func ServiceChatlistNewsY(userId int64) map[int][]map[string]interface{} {
	chatList := models.ObtainUserChatList(userId)
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
				item[v.UId] = append(item[v.UId], personMap)
			}
		}
	}
	//content := make(map[int][]map[string]interface{})
	//for key, _ := range item {
	//
	//	if key != int(userId) {
	//		data := models.ChatLastPieceData(int64(userId), key, 2)
	//		if key < 10 {
	//			data = models.ChatLastPieceData(int64(userId), key, 1)
	//		}
	//
	//		value := map[string]interface{}{}
	//		if strings.Contains(data.Content, "<img") {
	//			value["Content"] = "[图片]"
	//		} else if strings.Contains(data.Content, `{"id":`) {
	//			value["Content"] = "[订单]"
	//		} else {
	//			value["Content"] = data.Content
	//		}
	//		value["Count"] = 0
	//		if int(userId) == data.PId {
	//			value["Count"] = models.CountUserMessages(int64(userId), key)
	//		}
	//
	//		value["UID"] = key
	//		if key < 10 {
	//			value["UAvatar"] = "images/IN7gUqUPXXK2AGgepnGVk1fq5rVRZj7NqCSXO4NB.png"
	//			value["UName"] = "客服"
	//			value["Phone"] = ""
	//		} else {
	//			user := models.GetUserByFirstValue("name,avatar,phone", key)
	//			value["UAvatar"] = user.Avatar
	//			value["UName"] = user.Name
	//			value["Phone"] = user.Phone
	//		}
	//		value["State"] = data.State
	//		value["created_at"] = data.CreatedAt
	//		value["media"] = data.Media
	//		content[key] = append(content[key], value)
	//	}
	//}
	return item
}

// 自定义排序函数
//func arraySort(data []map[string]interface{}, keys string) []map[string]interface{} {
//	timestamps := make([]int64, len(data))
//
//	for key, row := range data {
//		if value, ok := row[keys].(string); ok {
//			timestamp, err := strconv.ParseInt(value, 10, 64)
//			if err == nil {
//				timestamps[key] = timestamp
//			}
//		}
//	}
//
//	sort.Slice(data, func(i, j int) bool {
//		return timestamps[i] > timestamps[j]
//	})
//
//	return data
//}

//type Data struct {
//	Count     int
//	UID       int
//	Avatar    string
//	Name      string
//	Phone     string
//	State     string
//	CreatedAt time.Time
//	Media     string
//}

//func structToMap(s interface{}) (map[string]interface{}, error) {
//	result := make(map[string]interface{})
//
//	val := reflect.ValueOf(s)
//	if val.Kind() == reflect.Ptr {
//		val = val.Elem()
//	}
//
//	if val.Kind() != reflect.Struct {
//		return nil, fmt.Errorf("not a struct")
//	}
//
//	typ := val.Type()
//	for i := 0; i < val.NumField(); i++ {
//		field := typ.Field(i)
//		if !val.Field(i).CanInterface() {
//			continue
//		}
//		name := field.Tag.Get("json")
//		if name == "" {
//			name = field.Name
//		}
//		result[name] = val.Field(i).Interface()
//	}
//
//	return result, nil
//}
