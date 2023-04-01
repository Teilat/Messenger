package cache

import (
	"Messenger/internal/database"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func (c *cache) CreateMessage(msg *database.Message) error {
	if _, ok := c.message[msg.Id]; ok {
		return fmt.Errorf("msg with id:%s already exist", msg.Id.String())
	}
	err := validateMessage(msg)
	if err != nil {
		return err
	}
	c.message[msg.Id] = msg

	c.updateChan <- database.UpdateMessage{Message: []*database.Message{msg}}
	return nil
}

func (c *cache) UpdateMessage(msg *database.Message) error {
	if _, ok := c.message[msg.Id]; !ok {
		return fmt.Errorf("msg with id:%s does not exist", msg.Id.String())
	}
	err := validateMessage(msg)
	if err != nil {
		return err
	}
	c.message[msg.Id] = msg

	// send updates in chan
	c.updateChan <- database.UpdateMessage{Message: []*database.Message{msg}}

	return nil
}

func (c *cache) DeleteMessage(id uuid.UUID) error {
	msg, ok := c.message[id]
	if !ok {
		return fmt.Errorf("msg with id:%s does not exist", id.String())
	}

	// send updates in chan
	c.updateChan <- database.UpdateMessage{Message: []*database.Message{msg}}
	return nil
}

func (c *cache) Message(id uuid.UUID) (*database.Message, bool) {
	msg, ok := c.message[id]
	if !ok {
		return nil, false
	}
	if msg.DeletedForAll {
		return nil, false
	}
	return msg, ok
}

func validateMessage(msg *database.Message) error {
	var err error
	if msg.Id == uuid.Nil {
		err = errors.Join(err, fmt.Errorf("id nil or emprty"))
	}
	if msg.ChatId == uuid.Nil {
		err = errors.Join(err, fmt.Errorf("chat id nil or emprty"))
	}
	if msg.Username == "" {
		err = errors.Join(err, fmt.Errorf("username emprty"))
	}
	return err
}
