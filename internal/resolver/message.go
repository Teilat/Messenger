package resolver

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"sort"
	"strconv"
	"time"
)

func (r Resolver) CreateWsMessage(msg models.SendMessage, chatId, userId string) error {
	chats := r.GetUserChats(userId)
	sort.Slice(chats, func(i, j int) bool {
		return chats[i].CreatedAt.Before(chats[j].CreatedAt)
	})
	chat, err := strconv.Atoi(chatId)
	if err != nil {
		return err
	}
	res := r.Db.Create(&database.Message{
		Text:      msg.Text,
		CreatedAt: time.Now(),
		Username:  userId,
		ChatId:    chats[chat].Id,
	})
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r Resolver) GetWsMessages(payload models.GetMessages, chatId, userId string) ([]database.Message, error) {
	chats := r.GetUserChats(userId)
	sort.Slice(chats, func(i, j int) bool {
		return chats[i].CreatedAt.Before(chats[j].CreatedAt)
	})
	chat, err := strconv.Atoi(chatId)
	if err != nil {
		return nil, err
	}
	return chats[chat].Messages[payload.Offset : payload.Offset+payload.Limit], nil
}
