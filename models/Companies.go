package models

import (
	"time"
)

type Companies struct {
	ID                  uint
	Number              string    `json:"number"`
	Name                string    `json:"name"`
	Logo                string    `json:"logo"`
	Status              int       `json:"status"`
	Attest              int       `json:"attest"`
	Address             string    `json:"address"`
	Phone               string    `json:"phone"`
	QrCode              string    `json:"qr_code"`
	UserId              int       `json:"user_id"`
	UserName            string    `json:"user_name"`
	TypeId              int       `json:"type_id"`
	RegionId            int       `json:"region_id"`
	Website             string    `json:"website"`
	Keyword             string    `json:"keyword"`
	Description         string    `json:"description"`
	Introduction        string    `json:"introduction"`
	Label               string    `json:"label"`
	MainItems           string    `json:"main_items"`
	Cover               string    `json:"cover"`
	License             int       `json:"license"`
	AccessNum           int       `json:"access_num"`
	GzNum               int       `json:"gz_num"`
	TicketNum           string    `json:"ticket_num"`
	Gallery             string    `json:"gallery"`
	RegisteredAddress   string    `json:"registered_address"`
	RegisteredCapital   string    `json:"registered_capital"`
	RegistrationNumber  string    `json:"registration_number"`
	LegalRepresentative string    `json:"legal_representative"`
	MainProducts        string    `json:"main_products"`
	SetUpTheTime        string    `json:"set_up_the_time"`
	CreditCode          string    `json:"credit_code"`
	Theme               int       `json:"theme"`
	Auth                int       `json:"auth"`
	PhoneAuth           int       `json:"phone_auth"`
	IdCardAuth          int       `json:"idCard_auth"`
	Vip                 int       `json:"vip"`
	IdNumber            string    `json:"id_number"`
	IdNumberImgZ        string    `json:"id_number_img_z"`
	IdNumberImgF        string    `json:"id_number_img_f"`
	LRphone             string    `json:"L_R_phone"`
	Likes               int       `json:"likes"`
	FreeLabel           string    `json:"free_label"`
	Banner              string    `json:"banner"`
	Longitude           string    `json:"longitude"`
	Latitude            string    `json:"latitude"`
	Classify            string    `json:"classify"`
	Route               string    `json:"route"`
	TradingAlliance     string    `json:"trading_alliance"`
	TradingAllianceType string    `json:"trading_alliance_type"`
	Liveness            int       `json:"liveness"`
	CreatedAt           time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt           time.Time `time_format:"2006-01-02 15:04:05"`
}
