package models

import (
	"time"
)

type ThinkChat struct {
	ID        uint
	Uid       int       `json:"u_id"`
	Uname     string    `json:"u_name"`
	Pid       int       `json:"p_id"`
	Pname     string    `json:"p_name"`
	Content   string    `json:"content"`
	State     int       `json:"state"`
	Media     int       `json:"media"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	Avatar    string    `json:"avatar"`
	Name      string    `json:"name"`
}

func (ThinkChat) TableName() string {
	return "think_chat"
}

func ObtainUserChatList(UserIdS int64) []ThinkChat {
	var Order []ThinkChat
	ChatDB.
		Where("think_chat.p_id = ? OR think_chat.u_id = ?", UserIdS, UserIdS).
		Joins("left join users ON users.id = think_chat.u_id").
		Joins("left join companies ON users.company_id = companies.id").
		Select(
			"think_chat.id," +
				"think_chat.u_id," +
				"think_chat.u_name," +
				"think_chat.p_id," +
				"think_chat.p_name," +
				"think_chat.content," +
				"think_chat.state," +
				"think_chat.created_at," +
				"think_chat.media," +
				"users.avatar," +
				"companies.name").
		Find(&Order)
	return Order
}
