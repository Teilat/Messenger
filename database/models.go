package database

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username    string    `gorm:"index:,unique"` // starts with @
	Name        string
	Description string
	PwHash      string
	CreatedAt   time.Time
	LastOnline  time.Time
	Image       []byte
	Chats       []*Chat `gorm:"many2many:user_chats;"` // many to many
}

type Chat struct {
	Id           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name         string
	CreationDate time.Time
	Users        []*User   `gorm:"many2many:user_chats;"` // many to many
	Messages     []Message // one to many
}

type Message struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Text      string
	CreatedAt time.Time
	EditedAt  time.Time
	DeletedAt time.Time
	ChatId    uuid.UUID // one to many chat id
}
