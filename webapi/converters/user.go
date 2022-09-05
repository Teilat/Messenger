package converters

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"strconv"
)

func UserToApiUser(user database.User) models.User {
	return models.User{
		Username:   user.Username,
		Nickname:   user.Name,
		Bio:        user.Bio,
		Phone:      user.Phone,
		CreatedAt:  strconv.FormatInt(user.CreatedAt.Unix(), 10),
		LastOnline: strconv.FormatInt(user.CreatedAt.Unix(), 10),
		Chats:      chatsToApiUserChats(user.Chats),
	}
}

func chatsToApiUserChats(chats []*database.Chat) []models.Chat {
	res := []models.Chat{}
	for _, chat := range chats {
		res = append(res, models.Chat{
			Name:      chat.Name,
			CreatedAt: strconv.FormatInt(chat.CreatedAt.Unix(), 10),
			LastMessage: models.Message{
				Text:      "",
				CreatedAt: "",
				User:      false,
			},
		})
	}
	return res
}
