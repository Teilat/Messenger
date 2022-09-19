package converters

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"sort"
)

const DefaultTimeFormat = "15:04"

func UserToApiUser(user *database.User, chats []*database.Chat) models.User {
	return models.User{
		Username:   user.Username,
		Nickname:   user.Name,
		Bio:        user.Bio,
		Phone:      user.Phone,
		CreatedAt:  user.CreatedAt.Format(DefaultTimeFormat),
		LastOnline: user.LastOnline.Format(DefaultTimeFormat),
		Chats:      chatsToApiUserChats(chats),
	}
}

func chatsToApiUserChats(chats []*database.Chat) []models.Chat {
	var res []models.Chat
	sort.Slice(chats, func(i, j int) bool {
		return chats[i].CreatedAt.Before(chats[j].CreatedAt)
	})

	for i, chat := range chats {
		sort.Slice(chat.Messages, func(i, j int) bool {
			return chat.Messages[i].CreatedAt.After(chat.Messages[j].CreatedAt)
		})

		res = append(res, models.Chat{
			Id:        uint32(i),
			Name:      chat.Name,
			CreatedAt: chat.CreatedAt.Format(DefaultTimeFormat),
			LastMessage: models.Message{
				Text:      chat.Messages[0].Text,
				CreatedAt: chat.Messages[0].CreatedAt.Format(DefaultTimeFormat),
				User:      "Admin",
			},
		})
	}
	return res
}
