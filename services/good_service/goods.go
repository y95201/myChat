/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-12-19 10:09:37
 * @LastEditors: Y95201
 * @LastEditTime: 2022-12-20 10:49:21
 */
package good_service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GoodsList(c *gin.Context) {
	//session.ClearAuthSession(c)
	c.Redirect(http.StatusFound, "/")
	return
}
