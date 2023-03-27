package cache

import (
	"Messenger/internal/database"
	"Messenger/internal/logger"
	"testing"
)

// full test for CRUD
func Test_Chat(t *testing.T) {
	c := NewCache(logger.NewLogger("[Test Cache]"))

	if err := c.CreateChat(&database.Chat{
		Name:     "",
		Users:    nil,
		Messages: nil,
	}); err != nil {
		t.Errorf("CreateChat() error = %v", err)
	}

}
