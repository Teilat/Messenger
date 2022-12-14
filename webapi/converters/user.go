package converters

import (
	"Messenger/database"
	"Messenger/webapi/models"
)

const DefaultTimeFormat = "15:04"

func UserToApiUser(user *database.User) models.User {
	return models.User{
		Username:   user.Username,
		Nickname:   user.Name,
		Bio:        user.Bio,
		Phone:      user.Phone,
		CreatedAt:  user.CreatedAt.Format(DefaultTimeFormat),
		LastOnline: user.LastOnline.Format(DefaultTimeFormat),
		Image:      user.Image,
	}
}

func chatUsersToApiChatUsers(users []*database.User) []string {
	var res []string
	for _, user := range users {
		res = append(res, user.Username)
	}
	return res
}
