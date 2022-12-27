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
	"myChat/services/good_service"
)

func Orderlist(c *gin.Context) {
	good_service.UserOrderList(c)
}
