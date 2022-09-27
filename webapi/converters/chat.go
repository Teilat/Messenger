package converters

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"sort"
	"strconv"
)

func ChatsToApiChats(chats []*database.Chat) []*models.Chat {
	res := make([]*models.Chat, 0)
	sort.Slice(chats, func(i, j int) bool {
		return chats[i].CreatedAt.Before(chats[j].CreatedAt)
	})
	for i, chat := range chats {
		res = append(res, &models.Chat{
			Id:          uint32(i),
			Name:        chat.Name,
			CreatedAt:   strconv.FormatInt(chat.CreatedAt.Unix(), 10),
			LastMessage: models.Message{},
		})
	}
	return res
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
		if len(chat.Messages) < 1 || chat.Messages == nil {
			chat.Messages = []database.Message{{}}
		}
		res = append(res, models.Chat{
			Id:        uint32(i),
			Name:      chat.Name,
			CreatedAt: chat.CreatedAt.Format(DefaultTimeFormat),
			LastMessage: models.Message{
				Text:      chat.Messages[0].Text,
				CreatedAt: chat.Messages[0].CreatedAt.Format(DefaultTimeFormat),
				UserId:    chat.Messages[0].Username,
			},
			Users: chatUsersToApiChatUsers(chat.Users),
		})
	}
	return res
}

func ChatToApiChat(chat *database.Chat) *models.Chat {
	return &models.Chat{
		Id:          0,
		Name:        chat.Name,
		CreatedAt:   chat.CreatedAt.Format(DefaultTimeFormat),
		LastMessage: models.Message{},
	}
}
