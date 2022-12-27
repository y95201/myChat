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
	"myChat/services/quick_service"
)

func Quicklist(c *gin.Context) {
	quick_service.List(c)
}
func Quickadd(c *gin.Context) {
	quick_service.Insert(c)
}

func QuickDell(c *gin.Context) {
	quick_service.Destroy(c)
}
