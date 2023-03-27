package database

import (
	"time"
)

type User struct {
	BaseModel
	Username   string `gorm:"index:,unique"` // starts with @
	Name       string
	Bio        string
	Phone      string
	PwHash     string
	CreatedAt  time.Time `gorm:"default:(now() at time zone 'msk')"`
	LastOnline time.Time `gorm:"default:(now() at time zone 'msk')"`
	Image      []byte
	Chats      []*Chat   `gorm:"many2many:user_chats;"`                   // many to many
	Messages   []Message `gorm:"foreignKey:Username;references:Username"` // one to many
}
