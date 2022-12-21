package models

import (
	"time"
)

type ThinkChat struct {
	ID        uint
	Uid       int       `json:"u_id"`
	Uname     string    `json:"u_name"`
	Pid       int       `json:"avatar_id"`
	Pname     string    `json:"p_name"`
	Content   string    `json:"content"`
	State     int       `json:"state"`
	Media     int       `json:"media"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}
