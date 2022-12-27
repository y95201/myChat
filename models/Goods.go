package models

import (
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
type Goods struct {
	ID        uint
	Contract  string    `json:"contract"`
	Count     int       `json:"count"`
	OrderType int       `json:"order_type"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}

func GetGoodsBylist(UserIdS int) []Goods {

	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)

	var g []Goods
	ChatDB.Select("any_value(goods.id) as id",
		"any_value(goods.note) as contract",
		"COUNT(goods.note) AS count",
		"any_value(goods.created_at) as created_at",
		"any_value(goods.type) as order_type").
		Where("goods.user_id = ? "+
			"AND goods.state = ? "+
			"AND goods.number != ? "+
			"AND goods.type = ? "+
			"AND goods.created_at > ?", UserIdS, "1", "0", "2", yesTime).
		Group("contract").
		Order("created_at desc").
		Find(&g)

	return g
}
