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
