package resolver

import (
	"Messenger/internal/cache"
	"Messenger/internal/database"
	"Messenger/webapi/models"
	"fmt"
	"github.com/google/uuid"
	"sort"
	"strconv"
	"time"
)

func (r Resolver) CreateChat(chat models.AddChat) (*database.Chat, error) {
	id := uuid.New()
	err := r.Cache.CreateChat(&database.Chat{
		Id:        id,
		Name:      chat.Name,
		CreatedAt: time.Now(),
		Users:     makeUsersForChat(r.Cache, chat.Users),
		Messages:  []database.Message{},
	})
	if err != nil {
		return nil, err
	}

	res, ok := r.Cache.Chat(id)
	if !ok {
		return nil, fmt.Errorf("falied to get chat")
	}
	return res, nil
}

func (r Resolver) GetUserChats(userId string) []*database.Chat {
	id, _ := uuid.Parse(userId)
	user, ok := r.Cache.User(id)
	if !ok {
		r.Logger.Warning("User with id %s not found", userId)
		return nil
	}

	chats := r.Cache.ChatsByUser(user)

	// Sort chats by creation date TODO:sort by last msg date
	sort.Slice(chats, func(i, j int) bool {
		return chats[i].CreatedAt.Before(chats[j].CreatedAt)
	})
	// Sort messages in chat by creation date
	for _, chat := range chats {
		sort.Slice(chat.Messages, func(i, j int) bool {
			return chat.Messages[i].CreatedAt.Before(chat.Messages[j].CreatedAt)
		})
	}
	return chats
}

func (r Resolver) ChatIdToUUID(chatId string, userId string) uuid.UUID {
	chat, err := r.chatFromChatId(userId, chatId)
	if err != nil {
		return uuid.UUID{}
	}
	return chat.Id
}

func (r Resolver) chatFromChatId(userId, chatId string) (*database.Chat, error) {
	chats := r.GetUserChats(userId)
	chat, err := strconv.Atoi(chatId)
	if err != nil {
		return nil, err
	}
	return chats[chat], nil
}

func makeUsersForChat(c cache.Cache, usernames []string) []*database.User {
	return c.UsersByNames(usernames)
}
