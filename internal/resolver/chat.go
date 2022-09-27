package resolver

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"gorm.io/gorm"
	"time"
)

func (r Resolver) CreateChat(chat models.AddChat) (*database.Chat, error) {
	res := r.Db.Create(&database.Chat{
		Name:      chat.Name,
		CreatedAt: time.Now(),
		Users:     makeUsersForChat(r.Db, chat.Users),
		Messages:  nil,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Statement.Model.(*database.Chat), nil
}

func (r Resolver) GetUserChats(userId string) []*database.Chat {
	var chats []*database.Chat
	var res []*database.Chat
	r.Db.
		Preload("Users").
		Preload("Messages").
		Find(&chats)

	// TODO сделать фильтауию через запрос
	for _, chat := range chats {
		for _, user := range chat.Users {
			if user.Username == userId {
				res = append(res, chat)
			}
		}
	}
	return res
}

func makeUsersForChat(db *gorm.DB, usernames []string) []*database.User {
	user := make([]*database.User, 0)
	db.Preload("Chats").Where("username IN ?", usernames).Find(&user)
	return user
}
