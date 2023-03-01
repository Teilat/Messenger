package cache

import (
	"Messenger/internal/database"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (c *cache) CreateMessage(msg *database.Message) error {
	if _, ok := c.message[msg.Id]; ok {
		return fmt.Errorf("msg with id:%d already exist", msg.Id)
	}
	msg.CreatedAt = time.Now()
	c.message[msg.Id] = msg

	c.updateChan <- UpdateMessage{Message: []*database.Message{msg}}
	return nil
}

func (c *cache) UpdateMessage(msg *database.Message) error {
	if _, ok := c.message[msg.Id]; !ok {
		return fmt.Errorf("msg with id:%d does not exist", msg.Id)
	}

	msg.EditedAt = time.Now()
	c.message[msg.Id] = msg

	// send updates in chan
	c.updateChan <- UpdateMessage{Message: []*database.Message{msg}}

	return nil
}

func (c *cache) DeleteMessage(id uuid.UUID, deletedForAll bool) error {
	msg, ok := c.message[id]
	if ok {
		return fmt.Errorf("msg with id:%d does not exist", id)
	}

	// update msg deleted at
	msg.DeletedAt = time.Now()
	msg.DeletedForAll = deletedForAll

	// send updates in chan
	c.updateChan <- UpdateMessage{Message: []*database.Message{msg}}
	return nil
}

func (c *cache) Message(id uuid.UUID) (*database.Message, bool) {
	msg, ok := c.message[id]
	return msg, ok
}
