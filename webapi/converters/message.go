package converters

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"sort"
)

func MessagesToWsMessages(msg []database.Message) []models.Message {
	res := []models.Message{}

	for _, message := range msg {
		res = append(res, models.Message{
			Id:         message.IntId,
			Text:       message.Text,
			CreatedAt:  message.CreatedAt.Format(DefaultTimeFormat),
			UserId:     message.Username,
			ResponseId: *message.ResponseId,
		})
	}
	return res
}

func messagesToFirstApiMessages(messages []database.Message) []models.Message {
	res := make([]models.Message, 0)
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].CreatedAt.After(messages[j].CreatedAt)
	})
	for i, message := range messages {
		if i >= 30 {
			break
		}
		res = append(res, models.Message{
			Id:        message.IntId,
			Text:      message.Text,
			CreatedAt: message.CreatedAt.Format(DefaultTimeFormat),
			UserId:    message.Username,
		})
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
