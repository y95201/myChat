package models

import (
	"errors"

	"gorm.io/gorm"
)

type ThinkChatBinding struct {
	ID   uint
	Uid  int `json:"u_id"`
	Ufd  int `json:"u_fd"`
	Pid  int `json:"p_id"`
	Type int `json:"type"`
}

func (m *ThinkChatBinding) ThinkChatBinDingTableName() string {
	return "think_chat_binding"
}

// 查询单个数据判断
func ThinkChatQuickExists(u_id int, p_id int) bool {
	var u ThinkChatBinding
	result := ChatDB.Where("u_id = ?", u_id).Where("p_id = ?", p_id).First(&u)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}
func ThinkChatQuickDelete(u_id int) {
	var u ThinkChatBinding
	ChatDB.Where("u_id = ?", u_id).Delete(&u)
}
func ThinkChatQuickSave(data *ThinkChatBinding) {
	ChatDB.Save(data)
}
