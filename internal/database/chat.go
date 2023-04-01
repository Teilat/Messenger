package database

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string
	CreatedAt time.Time `gorm:"default:(now() at time zone 'msk')"`
	Users     []*User   `gorm:"many2many:user_chats;"` // many to many
	Messages  []Message `gorm:"foreignKey:ChatId"`     // one to many
}

func (c *Chat) GetId() uuid.UUID {
	if c.Id != uuid.Nil {
		return c.Id
	}
	return uuid.Nil
}

func (c *Chat) GetMessageById(id uint32) (*Message, error) {
	for _, msg := range c.Messages {
		if msg.IntId == id {
			return &msg, nil
		}
	}
	return nil, fmt.Errorf("no message found with id:%d", id)
}
