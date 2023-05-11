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

func (m *ThinkChatQuick) ThinkChatQuickTableName() string {
	return "think_chat_quick"
}
func GetChatQuickByList(field, value string) []ThinkChatQuick {

	var articles []ThinkChatQuick

	if field == "userId" {
		ChatDB.Where(field+" = ?", value).Find(&articles)
	}

	return articles
}
func SetChatQuickByCreate(value interface{}) ThinkChatQuick {
	var m ThinkChatQuick

	m.UserId = value.(map[string]interface{})["user_id"].(int)
	m.State = value.(map[string]interface{})["state"].(int)
	m.Content = value.(map[string]interface{})["content"].(string)
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	ChatDB.Create(&m)
	return m
}

func SetChatQuickByDelete(field, value string) []ThinkChatQuick {

	var articles []ThinkChatQuick

	if field == "id" {
		ChatDB.Where(field+" = ?", value).Delete(&articles)
	}

	return articles
}
