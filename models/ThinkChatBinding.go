package models

type ThinkChatBinding struct {
	ID   uint
	Uid  int `json:"u_id"`
	Ufd  int `json:"u_fd"`
	Pid  int `json:"p_id"`
	Type int `json:"type"`
}
