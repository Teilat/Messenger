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

type Chat struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string
	CreatedAt time.Time `gorm:"default:(now() at time zone 'msk')"`
	Users     []*User   `gorm:"many2many:user_chats;"` // many to many
	Messages  []Message `gorm:"foreignKey:ChatId"`     // one to many
}

type Message struct {
	Id            uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	IntId         uint32
	Text          string
	CreatedAt     time.Time `gorm:"default:(now() at time zone 'msk')"`
	EditedAt      time.Time
	DeletedAt     time.Time
	DeletedForAll bool
	ResponseId    *uint32   `gorm:"foreignKey:IntId"`
	Username      string    // one to many username
	ChatId        uuid.UUID // one to many chat id
}
