package models

import (
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
	UsageTime   int     `json:"usage_time"`
	UserId      string  `json:"user_id"`
	OrderMoney  float32 `json:"order_money"`
}

func (Good) TableName() string { return "goods" }

func GetSellGoodsBylist(UserIdS int) []Goods {
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
	return sellGoods
}
func GetDetailedProductList(Contract string, wheres string) []Good {
	var list []Good
	where := map[string]interface{}{}
	where[wheres] = Contract
	ChatDB.
		Where(where).
		Joins("left join orders on orders.goods_id = goods.id").
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
	return list
}
func GetBuyGoodsBylist(UserIdS int) []Goods {
	var sellGoods []Goods
	ChatDB.
		Where("orders.user_id = ?", UserIdS).
		Where("goods.user_id != ?", "").
		Where("goods.type = ?", 4).
		Where("goods.state = ?", 1).
		Where("goods.number != ?", 0).
		Where("goods.created_at > ?", time.Now().AddDate(0, 0, -1)).
		Joins("left join orders on orders.goods_id = goods.id").
		Select("any_value(goods.id) as id," +
			"any_value(orders.contract) as contract," +
			"COUNT(orders.contract) AS count," +
			"any_value(goods.created_at) as created_at," +
			"any_value(goods.type) as order_type").
		Group("orders.contract").
		Order("created_at DESC").
		Find(&sellGoods)
	return sellGoods
}
func GetDepositOrderList(contract []string) []Goods {
	var depositGoods []Goods
	where := map[string]interface{}{}
	where["orders.contract"] = contract
	ChatDB.
		Where(where).
		Joins("left join orders on orders.goods_id = goods.id").
		Select("any_value(goods.id) as id," +
			"any_value(orders.contract) as contract," +
			"COUNT(orders.contract) AS count," +
			"any_value(goods.created_at) as created_at," +
			"any_value(goods.type) as order_type").
		Group("orders.contract").
		Order("created_at DESC").
		Find(&depositGoods)
	return depositGoods
}
