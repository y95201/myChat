package models

import (
	"time"
)

type ThinkChatUser struct {
	ID        uint
	Fid       int       `json:"f_id"`
	Uid       int       `json:"u_id"`
	IsUser    int       `json:"is_user"`
	Uname     string    `json:"u_name"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}
