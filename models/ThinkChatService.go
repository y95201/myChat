package models

import (
	"time"
)

type ThinkChatService struct {
	ID        uint
	Name      string    `json:"name"`
	Pass      string    `json:"pass"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}

func (m *ThinkChatService) TableName() string {
	return "think_chat_service"
}
func FindChatUserByField(field, value string) ThinkChatService {
	var u ThinkChatService

	if field == "name" {
		ChatDB.Where(field+" = ?", value).First(&u)
	}

	return u
}
