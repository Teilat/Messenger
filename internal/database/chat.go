package database

import (
	"fmt"
	"time"
)

type Chat struct {
	BaseModel
	Name      string
	CreatedAt time.Time `gorm:"default:(now() at time zone 'msk')"`
	Users     []*User   `gorm:"many2many:user_chats;"` // many to many
	Messages  []Message `gorm:"foreignKey:ChatId"`     // one to many
}

func (c *Chat) GetMessageById(id uint32) (*Message, error) {
	for _, msg := range c.Messages {
		if msg.IntId == id {
			return &msg, nil
		}
	}
	return nil, fmt.Errorf("no message found with id:%d", id)
}
