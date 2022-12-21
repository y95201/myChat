/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-12-20 14:00:12
 * @LastEditors: Y95201
 * @LastEditTime: 2022-12-20 14:43:48
 */
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// 已登录跳转room界面，多页面应该考虑放在中间件实现
	//userInfo := user_service.GetUserInfo(c)
	//if len(userInfo) > 0 {
	//	c.Redirect(http.StatusFound, "/home")
	//	return
	//}
	//
	//OnlineUserCount := primary.OnlineUserCount()

	c.HTML(http.StatusOK, "index.html", gin.H{})
}
