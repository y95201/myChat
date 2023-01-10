package models

import (
	"fmt"
	"reflect"
	"time"
)

type Goods struct {
	ID           int       `json:"id"`
	Contract     string    `json:"contract"`
	Count        int       `json:"count"`
	OrderType    string    `json:"order_type"`
	CreatedAt    time.Time `time_format:"2006-01-02 15:04:05"`
	Good         []Good    `json:"goods" gorm:"type:Good"`
	TotalWeight  float32   `json:"total_weight"`
	TotalAmount  float32   `json:"total_amount"`
	TotalDeposit float32   `json:"total_deposit"`
}

func (Goods) TableName() string { return "goods" }

type Good struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Texture     string  `json:"texture"`
	Spec        string  `json:"spec"`
	Price       string  `json:"price"`
	Number      string  `json:"number"`
	WeightTon   string  `json:"weight_ton"`
	Unit        string  `json:"unit"`
	TotalWeight float32 `json:"total_weight"`
	UsageTime   string  `json:"usage_time"`
	UserId      string  `json:"user_id"`
	OrderMoney  float32 `json:"order_money"`
}

func (Good) TableName() string { return "goods" }
func GetGoodsBylist(UserIdS int) []Goods {

	var sellGoods []Goods

	ChatDB.
		Where("goods.user_id = ?", UserIdS).
		Where("goods.state = ?", 1).
		Where("goods.number != ?", 0).
		Where("goods.type = ?", 2).
		Where("goods.created_at > ?", time.Now().AddDate(0, 0, -1)).
		Select("any_value(goods.id) as id, " +
			"any_value(goods.note) as contract, " +
			"COUNT(goods.note) AS count, " +
			"any_value(goods.created_at) as created_at, " +
			"any_value(goods.type) as order_type").
		Group("goods.note").
		Order("created_at DESC").
		Find(&sellGoods)

	var list []Good
	if len(sellGoods) > 0 {
		for i := range sellGoods {
			ChatDB.Where("goods.note = ?", sellGoods[i].Contract).
				Select("goods.id, " +
					"goods.name, " +
					"goods.texture, " +
					"goods.spec, " +
					"goods.price, " +
					"goods.number, " +
					"goods.weight_ton, " +
					"goods.unit, " +
					"goods.total_weight, " +
					"goods.usage_time, " +
					"goods.user_id , " +
					"round(goods.price * goods.total_weight , 2) as order_money").
				Find(&list)

			sellGoods[i].Good = list
			sellGoods[i].TotalWeight = ArraySum(list, "TotalWeight")
			sellGoods[i].TotalAmount = ArraySum(list, "OrderMoney")
			sellGoods[i].TotalDeposit = DepositAlgorithm(list)
		}
	}
	return sellGoods
}
func DepositAlgorithm(Orders []Good) float32 {
	var money float32 = 0

	if fmt.Sprintf("%f", Orders[0].TotalWeight) != "" {
		if Orders[0].UsageTime != "" {
			if IfUserId(Orders[0].UserId) == 1 {
				//orderMoney := int(math.Ceil(sum())) * 2
				//UsageMoney := strToInt(Orders[0].UsageTime) * 1
				//money = orderMoney + UsageMoney
				SumAndColumn(Orders, "TotalWeight")
			}
		} else {

			if IfUserId(Orders[0].UserId) == 1 {
				SumAndColumn(Orders, "TotalWeight")

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
func SumAndColumn(matrix []Good, fields string) float32 {
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
				fmt.Println(columns)
			}
		}
	}
	return columns
}

//
func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func ArraySum(array []Good, field string) float32 {
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
