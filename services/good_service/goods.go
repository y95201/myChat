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
	"reflect"
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
	//models.GetUserByFieldValue("id", IntUserId)
	//SellGoods := models.GetSellGoodsBylist(IntUserId)
	//if len(SellGoods) > 0 {
	//	for i := range SellGoods {
	//		var lists []models.Good
	//		lists = models.GetDetailedProductList(SellGoods[i].Contract)
	//		SellGoods[i].Good = lists
	//		SellGoods[i].TotalWeight = ArraySum(lists, "TotalWeight")
	//		SellGoods[i].TotalAmount = ArraySum(lists, "OrderMoney")
	//		SellGoods[i].TotalDeposit = DepositAlgorithm(lists)
	//	}
	//}

	BuyGoods := models.GetBuyGoodsBylist(IntUserId)
	if len(BuyGoods) > 0 {
		for i := range BuyGoods {
			var lists []models.Good
			lists = models.GetDetailedProductList(BuyGoods[i].Contract)
			BuyGoods[i].Good = lists
			BuyGoods[i].TotalWeight = ArraySum(lists, "TotalWeight")
			BuyGoods[i].TotalAmount = ArraySum(lists, "OrderMoney")
			BuyGoods[i].TotalDeposit = DepositAlgorithm(lists)
		}
	}

	// 将查询结果转换为二维数组
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  BuyGoods,
	})

	return
}
func ArraySum(array []models.Good, field string) float32 {
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

func DepositAlgorithm(Orders []models.Good) float32 {
	var money float32 = 0

	if fmt.Sprintf("%f", Orders[0].TotalWeight) != "" {
		if fmt.Sprintf("%v", Orders[0].UsageTime) != "0" {
			if IfUserId(Orders[0].UserId) == 1 {
				orderMoney := SumAndColumn(Orders, "TotalWeight") * 2
				usageMoney := Orders[0].UsageTime * MoneyEveryDay()
				money = orderMoney + float32(usageMoney)
			} else {
				orderMoney := SumAndColumn(Orders, "TotalWeight") * ReserveAmount()
				usageMoney := Orders[0].UsageTime * MoneyEveryDay()
				money = orderMoney + float32(usageMoney)
			}
		} else {
			if IfUserId(Orders[0].UserId) == 1 {
				money = SumAndColumn(Orders, "TotalWeight") * 2
			} else {
				money = SumAndColumn(Orders, "TotalWeight") * ReserveAmount()
			}
		}
	}
	//fmt.Println(matrix)
	return money
}
func ReserveAmount() float32 {
	return 200
}
func SumAndColumn(matrix []models.Good, fields string) float32 {
	var columns float32 = 0
	for _, row := range matrix {
		v := reflect.TypeOf(row)
		values := reflect.ValueOf(row)
		count := v.NumField()
		for i := 0; i < count; i++ {
			field := v.Field(i)
			value := values.Field(i)
			if field.Name == fields {
				iVal := value.Interface()
				num2 := iVal.(float32)
				columns += num2
			}
		}
	}
	return columns
}

//
func MoneyEveryDay() int {
	return 100
}

//
func IfUserId(userId string) int {
	switch userId {
	case "20790": // 18458315669
		return 1
	case "19197": // 18368150019
		return 1
	case "18858": // 18309852787
		return 1
	case "11422": // 18037123430
		return 1
	case "18582": // 18857132529
		return 1
	case "22708": // 18370026863 15958125229 22708
		return 1
	case "12179": // 17858642023
		return 1
	case "11385": // 15191660810
		return 1
	default:
		return 2
	}
}
