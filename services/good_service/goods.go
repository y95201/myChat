/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-12-19 10:09:37
 * @LastEditors: Y95201
 * @LastEditTime: 2022-12-20 10:49:21
 */
package good_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myChat/models"
	"net/http"
	"strconv"
)

type sortSons struct {
	models.Goods

	TotalWeight  int
	TotalAmount  int
	TotalDeposit int
}

func UserOrderList(c *gin.Context) {
	//2:定金、3:发现好物 4:行情锁价 5定金订货

	UserId := c.PostForm("user_id")
	if len(UserId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数为空"})
		return
	}
	IntUserId, _ := strconv.Atoi(UserId)
	models.GetUserByFieldValue("id", IntUserId)
	SellGoods := models.GetGoodsBylist(IntUserId)
	//result := make([][]string, len(SellGoods))
	for _, p := range SellGoods {
		fmt.Println(p)

	}
	// 将查询结果转换为二维数组

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  SellGoods,
	})

	return
}
func arraySum(array []Good, field string) float32 {
	sum := float32(0)
	for i := range array {
		switch field {
		case "TotalWeight":
			sum += array[i].TotalWeight
		case "OrderMoney":
			sum += array[i].OrderMoney
		}
	}
	return sum
}

//func ArraySumColumn(data interface{}, column string) (int)
//	// 定义一个 map 用于存储统计结果
//	m := make(map[string]int)
//
//	// 定义要统计的 key
//	key := column
//	// 使用嵌套的 for 循环遍历数组
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			for k, v := range data[i][j] {
//			// 如果是指定的 key，就将值相加
//			if k == key {
//				m[k] += v
//			}
//		}
//	}
//	return m[key]
//}
