package resolver

import (
	"Messenger/db"
	"Messenger/webapi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sort"
	"strconv"
	"time"
)

func (r Resolver) CreateChat(chat models.AddChat) (*db.Chat, error) {
	res := r.Db.Create(&db.Chat{
		Name:      chat.Name,
		CreatedAt: time.Now(),
		Users:     makeUsersForChat(r.Db, chat.Users),
		Messages:  nil,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Statement.Model.(*db.Chat), nil
}

func (r Resolver) GetUserChats(userId string) []*db.Chat {
	var chats []*db.Chat
	var res []*db.Chat
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
	sort.Slice(res, func(i, j int) bool {
		return res[i].CreatedAt.Before(res[j].CreatedAt)
	})
	for _, re := range res {
		sort.Slice(re.Messages, func(i, j int) bool {
			return re.Messages[i].CreatedAt.Before(re.Messages[j].CreatedAt)
		})
	}
	return res
}

func (r Resolver) ChatIdToUUID(chatId string, userId string) uuid.UUID {
	chat, err := r.chatFromChatId(userId, chatId)
	if err != nil {
		return uuid.UUID{}
	}
	return chat.Id
}

func (r Resolver) chatFromChatId(userId, chatId string) (*db.Chat, error) {
	chats := r.GetUserChats(userId)
	chat, err := strconv.Atoi(chatId)
	if err != nil {
		return nil, err
	}
	return chats[chat], nil
}

func makeUsersForChat(database *gorm.DB, usernames []string) []*db.User {
	user := make([]*db.User, 0)
	database.Preload("Chats").Where("username IN ?", usernames).Find(&user)
	return user
}
