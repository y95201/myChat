/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-12-19 10:09:37
 * @LastEditors: Y95201
 * @LastEditTime: 2022-12-20 10:49:21
 */
package quick_service

import (
	"github.com/gin-gonic/gin"
	"myChat/models"
	"net/http"
	"strconv"
)

func List(c *gin.Context) {
	UserId := c.PostForm("userId")
	if len(UserId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数为空"})
		return
	}
	list := models.GetChatQuickByList("user_id", UserId)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  list,
	})
	return
}

func Insert(c *gin.Context) {
	UserId := c.PostForm("userId")
	States := c.PostForm("state")
	Content := c.PostForm("content")
	if len(UserId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数为空"})
		return
	}
	if len(Content) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数为空"})
		return
	}
	if len(States) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数为空"})
		return
	}
	IntUserId, _ := strconv.Atoi(UserId)
	IntStates, _ := strconv.Atoi(States)
	models.SetChatQuickByCreate(map[string]interface{}{
		"user_id": IntUserId,
		"state":   IntStates,
		"content": Content,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
	})
	return
}
func Destroy(c *gin.Context) {

	Id := c.PostForm("Id")
	if len(Id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数为空"})
		return
	}
	models.SetChatQuickByDelete("id", Id)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
	})
	return
}
