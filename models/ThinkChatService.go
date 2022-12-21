package models

import (
	"time"
)

type ThinkChatService struct {
	ID        uint
	Uname     string    `json:"u_name"`
	Pass      string    `json:"pass"`
	Avatar    string    `json:"avatar"`
	Content   string    `json:"content"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}
