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
	State     string `json:"state" gorm:"column:state"`
	CreatedAt string `json:"created_at"`
	Media     string `json:"media"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
}

func ObtainUserChatList(UserIdS int64, keys int) []ThinkChat {
	var Order []ThinkChat
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
func ChatLastPieceData(UserIdS int64, key, keys int) ThinkChat {
	var Order ThinkChat
	if keys == 1 {
		ChatDB.
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
		ChatDB.
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
	//ThinkChatModel::create()->where('(think_chat.p_id = '.
	//	$user_id.' and think_chat.u_id = '.$key.') or
	//(think_chat.p_id = '.$key.' and think_chat.u_id = '.$user_id.')')->sum('state');
	return 1
}
