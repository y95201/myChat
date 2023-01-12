package models

import "time"

type Orders struct {
	Contract  string    `json:"contract"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UsageTime int       `json:"usage_time"`
}

func (Orders) TableName() string { return "orders" }
func GetDepositOrderInformation(UserIdS int) []Orders {
	var Order []Orders
	ChatDB.
		Where("goods.state = ?", 1).
		Where("goods.number != ?", 0).
		Where("goods.company_id = ?", 0).
		Where("goods.type = ?", 5).
		Where("orders.user_id = ?", UserIdS).
		Where("goods.deleted_at IS NULL").
		Joins("left join goods on orders.goods_id = goods.id").
		Select("DISTINCT `orders`.`contract`,`goods`.`created_at`,`goods`.`usage_time`").
		Find(&Order)
	return Order
}
