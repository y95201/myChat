package models

func (ThinkChat) TableName() string {
	return "think_chat"
}

type ThinkChat struct {
	ID        int    `json:"id" gorm:"column:id"`
	UId       int    `json:"u_id" gorm:"column:u_id"`
	UName     string `json:"u_name" gorm:"column:u_name"`
	PId       int    `json:"p_id" gorm:"column:p_id"`
	PName     string `json:"p_name" gorm:"column:p_name"`
	Content   string `json:"content" gorm:"column:content"`
	State     int    `json:"state" gorm:"column:state"`
	CreatedAt string `json:"created_at"`
	Media     int    `json:"media"`
	Url       string `json:"url"`
}

type ThinkChatData struct {
	ID        int    `json:"id" gorm:"column:id"`
	UId       int    `json:"u_id" gorm:"column:u_id"`
	UName     string `json:"u_name" gorm:"column:u_name"`
	PId       int    `json:"p_id" gorm:"column:p_id"`
	PName     string `json:"p_name" gorm:"column:p_name"`
	Content   string `json:"content" gorm:"column:content"`
	State     string `json:"state" gorm:"column:state"`
	CreatedAt string `json:"created_at"`
	Media     string `json:"media"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
}

func ObtainUserChatList(UserIdS int64, keys int) []ThinkChatData {
	var Order []ThinkChatData
	ChatDB := ChatDB.Table("think_chat").
		Joins("left join users ON users.id = think_chat.u_id").
		Joins("left join companies ON users.company_id = companies.id").
		Select(
			"think_chat.id," +
				"think_chat.u_id," +
				"think_chat.u_name," +
				"think_chat.p_id," +
				"think_chat.p_name," +
				"think_chat.content," +
				"think_chat.state," +
				"think_chat.created_at," +
				"think_chat.media," +
				"users.avatar," +
				"companies.name").
		Order("think_chat.created_at DESC")

	if keys == 1 {
		ChatDB = ChatDB.Where("think_chat.p_id = ? OR think_chat.u_id = ?", UserIdS, UserIdS)
	} else {
		ChatDB = ChatDB.Where("think_chat.p_id = 0 OR think_chat.p_id = ?", UserIdS)
	}
	ChatDB.Find(&Order)
	return Order
}

// 1为客服 2为用户
func ChatLastPieceData(UserIdS int64, key, keys int) ThinkChatData {
	var Order ThinkChatData
	if keys == 1 {
		ChatDB.Table("think_chat").
			Where("(think_chat.p_id = ? and think_chat.u_id < 10) "+
				"OR (think_chat.p_id < 10 and think_chat.u_id = ?)", UserIdS, UserIdS).
			Joins("left join users ON users.id = think_chat.u_id").
			Joins("left join companies ON users.company_id = companies.id").
			Select(
				"think_chat.id," +
					"think_chat.u_id," +
					"think_chat.u_name," +
					"think_chat.p_id," +
					"think_chat.p_name," +
					"think_chat.content," +
					"think_chat.state," +
					"think_chat.created_at," +
					"think_chat.media," +
					"users.avatar," +
					"companies.name").
			Order("think_chat.created_at DESC").
			Find(&Order)
	} else {
		ChatDB.Table("think_chat").
			Where("(think_chat.p_id = ? and think_chat.u_id = ?) "+
				"OR (think_chat.p_id = ? and think_chat.u_id = ?)", UserIdS, key, key, UserIdS).
			Joins("left join users ON users.id = think_chat.u_id").
			Joins("left join companies ON users.company_id = companies.id").
			Select(
				"think_chat.id," +
					"think_chat.u_id," +
					"think_chat.u_name," +
					"think_chat.p_id," +
					"think_chat.p_name," +
					"think_chat.content," +
					"think_chat.state," +
					"think_chat.created_at," +
					"think_chat.media," +
					"users.avatar," +
					"companies.name").
			Order("think_chat.created_at DESC").
			Find(&Order)
	}
	return Order
}
func CountUserMessages(UserIdS int64, key int) int {
	return 1
}
func GetLatestMessagesByUserId(userId int64) []ThinkChatData {
	var messages []ThinkChatData
	rawSQL := `
		WITH RankedChat AS (
			SELECT
				ROW_NUMBER() OVER (PARTITION BY think_chat.u_id ORDER BY think_chat.created_at DESC) AS RowNum,
				think_chat.id,
				think_chat.u_id,
				think_chat.u_name,
				think_chat.p_id,
				think_chat.p_name,
				think_chat.content,
				think_chat.state,
				think_chat.created_at,
				think_chat.media,
				think_chat.url,
				users.avatar,
				companies.name
			FROM
				think_chat
				LEFT JOIN users ON users.id = think_chat.u_id
				LEFT JOIN companies ON users.company_id = companies.id
			WHERE
				think_chat.p_id = 0 OR think_chat.p_id = ?
		)
		SELECT
			id,
			u_id,
			u_name,
			p_id,
			p_name,
			content,
			state,
			created_at,
			media,
			url,
			avatar,
			name
		FROM
			RankedChat
		WHERE
			RowNum = 1
		ORDER BY
			created_at DESC`

	// 执行原始 SQL 查询并将结果扫描到结构体中
	ChatDB.Raw(rawSQL, userId).Scan(&messages)
	return messages
}
func GetUserChatCount(userId int64, PId int64) int64 {
	var count int64
	ChatDB.Table("think_chat").
		Where("u_id = ? AND p_id = ? AND state = 1", userId, PId).
		Count(&count)
	return count
}
func ChatUserSaveMessageService(data *ThinkChat) {
	ChatDB.Save(data)
}

func PressThePrimaryKeyModify(Id int, updates map[string]interface{}) {
	var user ThinkChat
	ChatDB.Model(&user).Where("id = ?", Id).Updates(updates)
}

// 用户和用户聊天
func GetUserAndUserChatList(userid int, customer int) []ThinkChat {
	var Chat []ThinkChat
	ChatDB.Model(&Chat).
		Where("u_id = ?", userid).
		Where("p_id = ? OR u_id = ?", customer, customer).
		Where("p_id = ? OR u_id = ?", customer, customer).
		Where("p_id = 0").
		Find(&Chat)
	return Chat
}

// 用户和客服聊天
func GetUserAndServiceChatList(userid int) []ThinkChat {
	var chatList []ThinkChat
	ChatDB.Model(&ThinkChat{}).
		Where("(p_id = ? AND u_id IN (?)) OR (u_id = ? AND p_id IN (?))", userid, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, userid, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}).
		Find(&chatList)
	return chatList
}

// 客服和用户聊天
func GetServiceAndUserChatList(customer int) []ThinkChat {
	var chatList []ThinkChat
	ChatDB.Model(&ThinkChat{}).
		Where("(p_id = ? AND u_id IN (?)) OR (u_id = ? AND p_id IN (?))", customer, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, customer, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}).
		Find(&chatList)
	return chatList
}
