package resolver

import (
	"Messenger/internal/database"
	"Messenger/webapi/models"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sort"
	"strconv"
	"time"
)

func (r Resolver) CreateChat(chat models.AddChat) (*database.Chat, error) {
	resp := r.Db.Create(&database.Chat{
		Name:      chat.Name,
		CreatedAt: time.Now(),
		Users:     makeUsersForChat(r.Db, chat.Users),
		Messages:  nil,
	})
	if resp.Error != nil {
		return nil, resp.Error
	}

	res, ok := resp.Statement.Model.(*database.Chat)
	if !ok {
		return nil, fmt.Errorf("falied to type assert")
	}
	return res, nil
}

func (r Resolver) GetUserChats(userId string) []*database.Chat {
	var chats []*database.Chat
	var res []*database.Chat
	r.Db.
		Preload("Users").
		Preload("Messages").
		Find(&chats)

	// TODO сделать фильтрацию через запрос
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

func (r Resolver) chatFromChatId(userId, chatId string) (*database.Chat, error) {
	chats := r.GetUserChats(userId)
	chat, err := strconv.Atoi(chatId)
	if err != nil {
		return nil, err
	}
	return chats[chat], nil
}

func makeUsersForChat(db *gorm.DB, usernames []string) []*database.User {
	user := make([]*database.User, 0)
	db.Preload("Chats").Where("username IN ?", usernames).Find(&user)
	return user
}
