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
	userId := 4
	chatList := models.ObtainUserChatList(int64(userId))
	item := make(map[int][]map[string]interface{})
	for _, v := range chatList {

		personMap := map[string]interface{}{
			"Uid":       v.Uid,
			"Uname":     v.Uname,
			"Pid":       v.Pid,
			"Pname":     v.Pname,
			"Content":   v.Content,
			"State":     v.State,
			"Media":     v.Media,
			"CreatedAt": v.CreatedAt,
			"Avatar":    v.Avatar,
			"Name":      v.Name,
		}
		if _, ok := item[v.Uid]; !ok {

			if personMap["Uid"].(int) < 10 {
				personMap = map[string]interface{}{
					"Uid":       v.Uid,
					"Uname":     "钢信宝客服",
					"Pid":       v.Pid,
					"Pname":     v.Pname,
					"Content":   v.Content,
					"State":     v.State,
					"Media":     v.Media,
					"CreatedAt": v.CreatedAt,
					"Avatar":    "images/IN7gUqUPXXK2AGgepnGVk1fq5rVRZj7NqCSXO4NB.png",
					"Name":      v.Name,
				}

				item[0] = []map[string]interface{}{personMap}
			} else {

				//item[v.Uid] = []map[string]interface{}{personMap}
			}
		} else {
			//fmt.Println(v.Pid)
			//if v.Pid > 10 && v.Pid != userId {
			//	fmt.Println(personMap)
			//	item[v.Pid] = append(item[v.Pid], personMap)
			//}
		}
		//if _, ok := item[v.Uid]; !ok {
		//	item[v.Uid] = []map[string]interface{}{personMap} // 创建新的slice，将副本添加到其中
		//} else {
		//	item[v.Uid] = append(item[v.Uid], personMap) // 直接追加map副本
		//}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  item,
	})
	return
}

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
