package database

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username   string    `gorm:"index:,unique"` // starts with @
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
