package resolver

import (
	"Messenger/db"
	"Messenger/webapi/models"
	"time"
)

func (r Resolver) CreateWsMessage(msg models.SendMessage, chatId, userId string) (*db.Message, error) {
	chat, err := r.chatFromChatId(userId, chatId)
	if err != nil {
		return nil, err
	}
	res := r.Db.Create(&db.Message{
		IntId:     makeMessageId(chat),
		Text:      msg.Text,
		CreatedAt: time.Now(),
		Username:  userId,
		ChatId:    chat.Id,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Statement.Model.(*db.Message), nil
}

func (r Resolver) ReplyWsMessage(msg models.ReplyMessage, chatId, userId string) (*db.Message, error) {
	chat, err := r.chatFromChatId(userId, chatId)
	if err != nil {
		return nil, err
	}
	res := r.Db.Create(&db.Message{
		IntId:      makeMessageId(chat),
		Text:       msg.Text,
		CreatedAt:  time.Now(),
		Username:   userId,
		ChatId:     chat.Id,
		ResponseId: &msg.ReplyMessageId,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Statement.Model.(*db.Message), nil
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
	res := r.Db.Save(message)
	if res.Error != nil {
		return res.Error
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
	res := r.Db.Save(message)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r Resolver) GetWsMessages(payload models.GetMessages, chatId, userId string) ([]db.Message, error) {
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

func filterDeletedMessages(chat *db.Chat, userId string) {
	for i, message := range chat.Messages {
		if message.DeletedAt != time.Unix(0, 0) {
			if message.DeletedForAll || message.Username == userId {
				copy(chat.Messages[i:], chat.Messages[:i+1])
				continue
			}
		}
	}
}

func makeMessageId(chat *db.Chat) uint32 {
	return uint32(len(chat.Messages) + 1)
}
