package database

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type User struct {
	Id          pgtype.UUID    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username    pgtype.Varchar `gorm:"index:,unique"` // starts with @
	Name        pgtype.Varchar
	Description pgtype.Varchar
	PwHash      pgtype.Varchar
	CreatedAt   pgtype.Timestamp
	LastOnline  pgtype.Timestamp
	Image       pgtype.Bytea
	Chats       []*Chat `gorm:"many2many:user_chats;"` // many to many
}

type Chat struct {
	Id           pgtype.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name         pgtype.Name
	CreationDate pgtype.Date
	Users        []*User   `gorm:"many2many:user_chats;"` // many to many
	Messages     []Message // one to many
}

type Message struct {
	Id        pgtype.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Text      pgtype.Varchar
	CreatedAt pgtype.Timestamp
	EditedAt  pgtype.Timestamp
	DeletedAt gorm.DeletedAt
	ChatId    pgtype.UUID // one to many chat id
}
