package database

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	BaseModel
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
