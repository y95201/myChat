package models

import (
	"time"
)

type ThinkChatQuick struct {
	ID        uint
	UserId    int       `json:"user_id"`
	Content   string    `json:"content"`
	State     int       `json:"state"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}

func (m *ThinkChatQuick) TableName() string {
	return "think_chat_quick"
}
func GetChatQuickByList(field, value string) ThinkChatQuick {
	var u ThinkChatQuick

	if field == "name" {
		ChatDB.Where(field+" = ?", value).Find(&u)
	}

	return u
}
