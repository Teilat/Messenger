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
			Id:        uint32(i),
			Name:      chat.Name,
			CreatedAt: strconv.FormatInt(chat.CreatedAt.Unix(), 10),
			Messages:  messagesToFirstApiMessages(chat.Messages),
			Users:     chatUsersToApiChatUsers(chat.Users),
		})
	}
	return res
}

func ChatToApiChat(chat *database.Chat) *models.Chat {
	return &models.Chat{
		Id:        0,
		Name:      chat.Name,
		CreatedAt: chat.CreatedAt.Format(DefaultTimeFormat),
		Messages:  messagesToFirstApiMessages(chat.Messages),
		Users:     chatUsersToApiChatUsers(chat.Users),
	}
}
