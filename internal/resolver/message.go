package resolver

import (
	"Messenger/internal/database"
	"Messenger/webapi/models"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (r Resolver) CreateWsMessage(msg models.SendMessage, chatId, userId string) (*database.Message, error) {
	chat, err := r.chatFromChatId(userId, chatId)
	if err != nil {
		return nil, err
	}
	id := uuid.New()
	err = r.Cache.CreateMessage(&database.Message{
		Id:        id,
		IntId:     makeMessageId(chat),
		Text:      msg.Text,
		CreatedAt: time.Now(),
		Username:  userId,
		ChatId:    chat.Id,
	})
	if err != nil {
		return nil, err
	}
	res, ok := r.Cache.Message(id)

	if !ok {
		return nil, fmt.Errorf("falied to get message")
	}
	return res, nil
}

func (r Resolver) ReplyWsMessage(msg models.ReplyMessage, chatId, userId string) (*database.Message, error) {
	chat, err := r.chatFromChatId(userId, chatId)
	if err != nil {
		return nil, err
	}
	id := uuid.New()
	err = r.Cache.CreateMessage(&database.Message{
		Id:         id,
		IntId:      makeMessageId(chat),
		Text:       msg.Text,
		CreatedAt:  time.Now(),
		Username:   userId,
		ChatId:     chat.Id,
		ResponseId: &msg.ReplyMessageId,
	})
	if err != nil {
		return nil, err
	}

	res, ok := r.Cache.Message(id)
	if !ok {
		return nil, fmt.Errorf("falied to type assert")
	}
	return res, nil
}

func (r Resolver) EditWsMessage(payload models.EditMessage, chatId, userId string) error {
	chat, err := r.chatFromChatId(userId, chatId)
	if err != nil {
		return err
	}
	message, err := chat.GetMessageById(payload.MessageId)
	if err != nil {
		return err
	}

	message.Text = payload.NewText
	message.EditedAt = time.Now()
	err = r.Cache.UpdateMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func (r Resolver) DeleteWsMessage(payload models.DeleteMessage, chatId, userId string) error {
	chat, err := r.chatFromChatId(userId, chatId)
	if err != nil {
		return err
	}
	message, err := chat.GetMessageById(payload.MessageId)
	if err != nil {
		return err
	}

	message.DeletedAt = time.Now()
	err = r.Cache.DeleteMessage(message)
	if err != nil {
		return err
	}

	return nil
}

func (r Resolver) GetWsMessages(payload models.GetMessages, chatId, userId string) ([]database.Message, error) {
	if payload.Limit > 100 {
		payload.Limit = 100
	}
	chat, err := r.chatFromChatId(userId, chatId)
	if err != nil {
		return nil, err
	}

	filterDeletedMessages(chat, userId)

	if payload.Offset+payload.Limit > len(chat.Messages) {
		return chat.Messages[payload.Offset:], nil
	}
	return chat.Messages[payload.Offset : payload.Offset+payload.Limit], nil
}

func filterDeletedMessages(chat *database.Chat, userId string) {
	for i, message := range chat.Messages {
		if message.DeletedAt != time.Unix(0, 0) {
			if message.DeletedForAll || message.Username == userId {
				copy(chat.Messages[i:], chat.Messages[:i+1])
				continue
			}
		}
	}
}

func makeMessageId(chat *database.Chat) uint32 {
	return uint32(len(chat.Messages) + 1)
}
