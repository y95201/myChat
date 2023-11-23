package models

type ThinkChatUser struct {
	ID     uint   `gorm:"column:id"`
	Fid    int    `json:"fid" gorm:"column:fid"`
	Uid    int    `json:"u_id" gorm:"column:u_id"`
	Uname  string `json:"u_name" gorm:"column:u_name"`
	IsUser int    `json:"is_user" gorm:"column:is_user"`
}

func (ThinkChatUser) TableName() string {
	return "think_chat_user"
}

func GetListOfOnlineUsers() *[]ThinkChatUser {
	var u []ThinkChatUser

	ChatDB.Where("u_id > ?", 10).First(&u)

	return &u
}

func GetListOfOnlineCustomerService() *[]ThinkChatUser {
	var u []ThinkChatUser

	ChatDB.Where("u_id < ?", 10).First(&u)

	return &u
}

func GetChatUserService(userId int64) *[]ThinkChatUser {
	var u []ThinkChatUser

	ChatDB.Where("u_id = ?", userId).First(&u)

	return &u
}
func GetChatUserFirst(userId int64) ThinkChatUser {
	var u ThinkChatUser

	ChatDB.Where("u_id = ?", userId).First(&u)

	return u
}
func ChatUserSaveService(data *ThinkChatUser) {
	ChatDB.Save(data)
}

func ChatUserUpdatesService(userId int64, updates map[string]interface{}) {
	var user ThinkChatUser
	ChatDB.Model(&user).Where("u_id = ?", userId).Updates(updates)
}

func GetTheLastPieceData(customer int) int {
	var user ThinkChatUser

	ChatDB.Model(&user).Select("id").Where("u_id = ?", customer).Last(&user)
	return int(user.ID)
}
