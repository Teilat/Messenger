package converters

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"sort"
)

func MessagesToWsMessages(msg []database.Message) []models.Message {
	res := []models.Message{}

	for i, message := range msg {
		res = append(res, models.Message{
			Id:        uint32(i),
			Text:      message.Text,
			CreatedAt: message.CreatedAt.Format(DefaultTimeFormat),
			UserId:    message.Username,
		})
	}
	return res
}

func messagesToFirstApiMessages(messages []database.Message) []models.Message {
	res := make([]models.Message, 0)
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].CreatedAt.Before(messages[j].CreatedAt)
	})
	for i, message := range messages {
		if i > 30 {
			break
		}
		res = append(res, models.Message{
			Id:        uint32(i),
			Text:      message.Text,
			CreatedAt: message.CreatedAt.Format(DefaultTimeFormat),
			UserId:    message.Username,
		})
	}
	return res
}
