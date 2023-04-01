package cache

import (
	"Messenger/internal/database"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func (c *cache) CreateChat(chat *database.Chat) error {
	if _, ok := c.message[chat.Id]; ok {
		return fmt.Errorf("msg with id:%d already exist", chat.Id)
	}
	if err := c.validateChat(chat); err != nil {
		return err
	}
	c.chat[chat.Id] = chat

	// send updates in chan
	c.updateChan <- database.UpdateMessage{Chat: []*database.Chat{chat}}
	return nil
}

func (c *cache) ChatsByUser(user *database.User) []*database.Chat {
	res := make([]*database.Chat, 0)
	for _, chat := range c.chat {
		if containsUser(chat.Users, user) {
			res = append(res, chat)
		}
	}
	return res
}

func containsUser(users []*database.User, user *database.User) bool {
	for _, u := range users {
		if u.Id == user.Id {
			return true
		}
	}
	return false
}

func (c *cache) UpdateChat(updatedChat *database.Chat) error {
	// check if chat exist
	if _, ok := c.chat[updatedChat.Id]; !ok {
		return fmt.Errorf("chat with id:%d does not exist", updatedChat.Id)
	}
	if err := c.validateChat(updatedChat); err != nil {
		return err
	}
	c.chat[updatedChat.Id] = updatedChat

	// send updates in chan
	c.updateChan <- database.UpdateMessage{Chat: []*database.Chat{updatedChat}}
	return nil
}

func (c *cache) DeleteChat(id uuid.UUID) error {
	updMsg := database.UpdateMessage{}
	// check if chat exist
	deletedChat, ok := c.chat[id]
	if !ok {
		return fmt.Errorf("chat with id:%d does not exist", id)
	}
	// iterate over all chat users and remove chat from them
	for _, user := range deletedChat.Users {
		user.Chats = deleteChatFromSlice(user.Chats, deletedChat.Id)
	}

	// delete from cache
	delete(c.chat, id)

	// send updates in chan
	c.updateChan <- updMsg
	c.deleteChan <- database.DeleteMessage{Chat: []uuid.UUID{id}}
	return nil
}

func deleteChatFromSlice(slice []*database.Chat, id uuid.UUID) []*database.Chat {
	res := make([]*database.Chat, 0)
	for _, chat := range slice {
		if chat.Id == id {
			continue
		}
		res = append(res, chat)
	}
	return res
}

func (c *cache) Chat(id uuid.UUID) (*database.Chat, bool) {
	res, ok := c.chat[id]
	if !ok {
		return nil, false
	}
	return res, ok
}

func (c *cache) validateChat(chat *database.Chat) error {
	var err error
	if chat.Id == uuid.Nil {
		err = errors.Join(err, fmt.Errorf("id nil or emprty"))
	}
	if len(chat.Users) < 1 {
		err = errors.Join(err, fmt.Errorf("must have at least two users in chat"))
	}
	return err
}
