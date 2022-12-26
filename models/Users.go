package models

import (
	"time"
)

type Users struct {
	ID                  uint
	Phone               int       `json:"phone"`
	Nickname            string    `json:"nickname"`
	Name                string    `json:"name"`
	Password            string    `json:"password"`
	Age                 int       `json:"age"`
	Avatar              string    `json:"avatar"`
	Email               string    `json:"email"`
	Address             string    `json:"address"`
	Wechat              string    `json:"wechat"`
	Qq                  string    `json:"qq"`
	Token               string    `json:"token"`
	CompanyId           int       `json:"company_id"`
	Disable             int       `json:"disable"`
	Attest              int       `json:"attest"`
	Attention           string    `json:"attention"`
	TokenExpirationTime string    `json:"token_expiration_time"`
	MainRespFor         string    `json:"main_resp_for"`
	Sex                 int       `json:"sex"`
	Contact             string    `json:"contact"`
	Follow              int       `json:"follow"`
	Md5                 string    `json:"md5"`
	ContactStatus       int       `json:"contact_status"`
	Position            string    `json:"position"`
	DepartmentId        int       `json:"department_id"`
	PhoneShow           int       `json:"phone_show"`
	CompanyShow         int       `json:"company_show"`
	Autograph           string    `json:"autograph"`
	Background          string    `json:"background"`
	CardPhoto           string    `json:"card_photo"`
	Unionid             string    `json:"unionid"`
	Openid              string    `json:"openid"`
	AccountOpenid       string    `json:"account_openid"`
	RebateVip           string    `json:"rebate_vip"`
	RebateVipEndTime    string    `json:"rebate_vip_end_time"`
	AliAccount          string    `json:"Ali_account"`
	RealName            string    `json:"real_name"`
	CreatedAt           time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt           time.Time `time_format:"2006-01-02 15:04:05"`
	DpdatedAt           time.Time `time_format:"2006-01-02 15:04:05"`
}

//func AddUser(value interface{}) User {
//	var u User
//	u.Username = value.(map[string]interface{})["username"].(string)
//	u.Password = value.(map[string]interface{})["password"].(string)
//	u.AvatarId = value.(map[string]interface{})["avatar_id"].(string)
//	ChatDB.Create(&u)
//	return u
//}
//
func SaveAvatarId(u Users) Users {
	//u.AvatarId = AvatarId
	ChatDB.Save(&u)
	return u
}

func FindUserByField(field, value string) Users {
	var u Users

	if field == "id" || field == "username" {
		ChatDB.Where(field+" = ?", value).First(&u)
	}

	return u
}

//func GetOnlineUserList(uids []float64) []map[string]interface{} {
//	var results []map[string]interface{}
//	ChatDB.Where("id IN ?", uids).Find(&results)
//
//	return results
//}
