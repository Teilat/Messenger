package cache

import (
	"Messenger/internal/database"
	"fmt"
	"github.com/google/uuid"
)

func (c *cache) CreateChat(chat *database.Chat) error {
	// check if chat exist
	if _, ok := c.chat[chat.Id]; ok {
		return fmt.Errorf("chat with id:%d already exist", chat.Id)
	}

	c.chat[chat.Id] = chat

	// send updates in chan
	c.updateChan <- UpdateMessage{Chat: []*database.Chat{chat}}
	return nil
}

func (c *cache) UpdateChat(UpdatedChat *database.Chat) error {
	// check if chat exist
	if _, ok := c.chat[UpdatedChat.Id]; !ok {
		return fmt.Errorf("chat with id:%d does not exist", UpdatedChat.Id)
	}

	c.chat[UpdatedChat.Id] = UpdatedChat

	// send updates in chan
	c.updateChan <- UpdateMessage{Chat: []*database.Chat{UpdatedChat}}
	return nil
}

func (c *cache) DeleteChat(id uuid.UUID) error {
	updMsg := UpdateMessage{}
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
	c.deleteChan <- DeleteMessage{Chat: []uuid.UUID{id}}
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
	return res, ok
}
