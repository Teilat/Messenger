package db

import "fmt"

func (c *Chat) GetMessageById(id uint32) (*Message, error) {
	for _, msg := range c.Messages {
		if msg.IntId == id {
			return &msg, nil
		}
	}
	return nil, fmt.Errorf("no message found with id:%d", id)
}
