package resolver

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"strconv"
	"time"
)

func (r Resolver) CreateWsMessage(msg models.SendMessage, chatId, userId string) error {
	chats := r.GetUserChats(userId)
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

func (r Resolver) ReplyWsMessage(msg models.ReplyMessage, chatId, userId string) error {
	return nil
}

func (r Resolver) EditWsMessage(payload models.EditMessage, chatId, userId string) error {
	return nil
}

func (r Resolver) DeleteWsMessage(payload models.DeleteMessage, chatId, userId string) error {
	return nil
}

func (r Resolver) GetWsMessages(payload models.GetMessages, chatId, userId string) ([]database.Message, error) {
	if payload.Limit > 100 {
		payload.Limit = 100
	}
	chats := r.GetUserChats(userId)
	chat, err := strconv.Atoi(chatId)
	if err != nil {
		return nil, err
	}
	if payload.Offset+payload.Limit > len(chats[chat].Messages) {
		return chats[chat].Messages[payload.Offset:], nil
	}
	return chats[chat].Messages[payload.Offset : payload.Offset+payload.Limit], nil
}
