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
	Chats      []*Chat   `gorm:"many2many:user_chats;"` // many to many
	Messages   []Message // one to many
}

type Chat struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string
	CreatedAt time.Time `gorm:"default:(now() at time zone 'msk')"`
	Users     []*User   `gorm:"many2many:user_chats;"` // many to many
	Messages  []Message // one to many
}

type Message struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Text      string
	CreatedAt time.Time `gorm:"default:(now() at time zone 'msk')"`
	EditedAt  time.Time
	DeletedAt time.Time
	UserId    uuid.UUID // one to many user id
	ChatId    uuid.UUID // one to many chat id
}
