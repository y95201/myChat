/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-12-19 10:09:37
 * @LastEditors: Y95201
 * @LastEditTime: 2022-12-20 10:49:21
 */
package user_service

import (
	"github.com/gin-gonic/gin"
	"myChat/models"
	"myChat/services/cache"
	"myChat/services/validator"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {

	username := c.PostForm("username")
	pwd := c.PostForm("password")

	var u validator.User

	u.Username = username
	u.Password = pwd

	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 5000, "msg": err.Error()})
		return
	}

	user := models.FindChatUserByField("name", username)
	userInfo := user
	md5Pwd := pwd
	//md5Pwd := helper.Md5Encrypt(pwd)

	if userInfo.ID > 0 {
		// json 用户存在
		// 验证密码
		if userInfo.Pass != md5Pwd {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "密码错误",
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg":  "用户不存在",
		})
		return
	}

	if userInfo.ID > 0 {
		cache.SaveAuthCache(string(strconv.Itoa(int(userInfo.ID))))
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  userInfo,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 5001,
			"msg":  "系统错误",
		})
		return
	}
}

//func GetUserInfo(c *gin.Context) map[string]interface{} {
//	return session.GetSessionUserInfo(c)
//}
//
func Logout(c *gin.Context) {
	//session.ClearAuthSession(c)
	c.Redirect(http.StatusFound, "/")
	return
}
