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
