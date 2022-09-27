package converters

import (
	"Messenger/database"
	"Messenger/webapi/models"
)

func MessagesToWsMessages(msg []database.Message) []models.Message {
	res := []models.Message{}

	for _, message := range msg {
		res = append(res, models.Message{
			Text:      message.Text,
			CreatedAt: message.CreatedAt.Format(DefaultTimeFormat),
			UserId:    message.Username,
		})
	}

	return res
}
