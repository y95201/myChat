/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-12-20 14:00:12
 * @LastEditors: y95201 957612196@qq.com
 * @LastEditTime: 2023-07-28 14:52:54
 */
package controller

import (
	"fmt"
	"myChat/models"
	"myChat/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
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
	ServiceChatlist(9)
	// fmt.Println(userListMessage)
	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 0,
	// 	"msg":  userListMessage,
	// })
	// return
}
func ServiceChatlist(userId int64) {
	chatList := models.GetLatestMessagesByUserId(userId)
	fmt.Println(chatList)
}
