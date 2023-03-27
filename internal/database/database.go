package database

import (
	"Messenger/internal/cache"
	"github.com/google/uuid"
)

type BaseModel struct {
	Id uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
}

func (m BaseModel) GetId() uuid.UUID {
	if m.Id != uuid.Nil {
		return m.Id
	}
	return uuid.Nil
}

func (db *Database) StartUpdateListener(chan cache.UpdateMessage, chan cache.DeleteMessage) {

}
