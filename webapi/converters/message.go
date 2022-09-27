package converters

import (
	"Messenger/database"
	"Messenger/webapi/models"
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
