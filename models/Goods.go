package models

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

//type Goods struct {
//	ID               uint
//	Name             int       `json:"name"`
//	Spec             string    `json:"spec"`
//	Price            string    `json:"price"`
//	Seckill          string    `json:"seckill"`
//	Deposit          int       `json:"deposit"`
//	Number           string    `json:"number"`
//	Texture          string    `json:"texture"`
//	SteelMill        string    `json:"steel_mill"`
//	Warehouse        string    `json:"warehouse"`
//	WeightTon        string    `json:"weight_ton"`
//	Quality          string    `json:"quality"`
//	Craft            int       `json:"craft"`
//	UserId           int       `json:"user_id"`
//	CompanyId        int       `json:"company_id"`
//	Freight          string    `json:"freight"`
//	Service          string    `json:"service"`
//	Type             string    `json:"type"`
//	Give             int       `json:"give"`
//	Follow           string    `json:"follow"`
//	Remarks          int       `json:"remarks"`
//	TypeId           string    `json:"type_id"`
//	TypeTId          int       `json:"type_t_id"`
//	Images           string    `json:"images"`
//	Label            int       `json:"label"`
//	Sales            int       `json:"sales"`
//	Sort             int       `json:"sort"`
//	BeforePrice      string    `json:"before_price"`
//	State            string    `json:"state"`
//	WarehouseAddress string    `json:"warehouse_address"`
//	RegionId         string    `json:"region_id"`
//	Unit             string    `json:"unit"`
//	Note             string    `json:"note"`
//	WeightOnePiece   string    `json:"weight_one_piece"`
//	InventoryId      string    `json:"inventory_id"`
//	TotalWeight      string    `json:"total_weight"`
//	WeightPrice      string    `json:"weight_price"`
//	Phone            string    `json:"phone"`
//	CompanysId       string    `json:"companyId"`
//	UsageTime        string    `json:"usage_time"`
//	CreatedAt        time.Time `time_format:"2006-01-02 15:04:05"`
//	UpdatedAt        time.Time `time_format:"2006-01-02 15:04:05"`
//	DpdatedAt        time.Time `time_format:"2006-01-02 15:04:05"`
//}

type Good struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Name        string    `json:"name"`
	Texture     string    `json:"texture"`
	Spec        string    `json:"spec"`
	Contract    string    `json:"contract"`
	Price       float32   `json:"price"`
	Number      int       `json:"number"`
	WeightTon   float32   `json:"weight_ton"`
	Unit        string    `json:"unit"`
	TotalWeight float32   `json:"total_weight"`
	UsageTime   time.Time `time_format:"2006-01-02 15:04:05"`
	CreatedAt   time.Time `time_format:"2006-01-02 15:04:05"`
	State       int       `json:"state"`
	Type        int       `json:"type"`
	Note        string    `json:"note"`
}

func GetGoodsBylist(UserIdS int) []Good {
	var sellGoods []Good
	ChatDB.Where("goods.user_id = ?", UserIdS).
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
		Order("goods.created_at DESC").
		Find(&sellGoods)
	if len(sellGoods) > 0 {
		for i := range sellGoods {
			var lists []Good
			ChatDB.Where("goods.note = ?", sellGoods[i].Contract).
				Select("goods.id, goods.name, goods.texture, goods.spec, goods.price, goods.number, goods.weight_ton, goods.unit, goods.total_weight, goods.usage_time, goods.user_id , round(goods.price * goods.total_weight , 2) as order_money").
				Find(&lists)

			sellGoods[i].Goods = lists
			sellGoods[i].TotalWeight = math.Round(arraySum(lists, "TotalWeight"), 3)
			sellGoods[i].TotalAmount = math.Round(arraySum(lists, "OrderMoney"), 2)
			sellGoods[i].TotalDeposit = this.DepositAlgorithm(lists)
		}
	}
	return sellGoods
}
func arraySum(array []Goods, field string) float32 {
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
func (this *Type) DepositAlgorithm(Orders [][]string) int {
	money := 0
	if Orders[0]["total_weight"] != "" {
		if Orders[0]["usage_time"] != "" {
			// new algorithm
			if this.IfUserId(Orders[0]["user_id"]) == 1 {
				orderMoney := int(math.Ceil(sum(column(Orders, "total_weight")))) * 2
				usageMoney := strToInt(Orders[0]["usage_time"]) * 1
				money = orderMoney + usageMoney
			} else {
				orderMoney := int(math.Ceil(sum(column(Orders, "total_weight")))) * this.ReserveAmount()
				usageMoney := strToInt(Orders[0]["usage_time"]) * this.MoneyEveryDay()
				money = orderMoney + usageMoney
			}
		} else {
			if this.IfUserId(Orders[0]["user_id"]) == 1 {
				money = int(math.Ceil(sum(column(Orders, "total_weight")))) * 2
			} else {
				money = int(math.Ceil(sum(column(Orders, "total_weight")))) * this.ReserveAmount()
			}
		}
	} else {
		return fmt.Errorf("订单重量为空")
	}
	return money
}

func (this *Type) ReserveAmount() int {
	return 200
}

func (this *Type) MoneyEveryDay() int {
	return 100
}

func (this *Type) IfUserId(userId string) int {
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

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func column(matrix [][]string, i int) []int {
	column := make([]int, len(matrix))
	for j, row := range matrix {
		column[j] = strToInt(row[i])
	}
	return column
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

//func ProductInformation(contract string) {
//	result := map[string]interface{}{}
//	ChatDB.Table("goods").
//		Where("note = ?", contract).
//		Select("goods.id", "goods.name", "goods.texture", "goods.spec",
//			"goods.price", "goods.number", "goods.weight_ton", "goods.unit",
//			"goods.total_weight", "goods.usage_time", "goods.user_id",
//			"round(goods.price * goods.total_weight , 2) as order_money").
//		Scan(&result)
//	return
//}
