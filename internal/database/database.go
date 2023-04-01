package database

import (
	"context"
	"github.com/google/uuid"
)

type UpdateMessage struct {
	User    []*User
	Message []*Message
	Chat    []*Chat
}

type DeleteMessage struct {
	User    []uuid.UUID
	Message []uuid.UUID
	Chat    []uuid.UUID
}

func (db *Database) StartUpdateListener(ctx context.Context, udp chan UpdateMessage, del chan DeleteMessage) {
	go db.runUpdateListener(ctx, udp, del)
}

func (db *Database) runUpdateListener(ctx context.Context, upd chan UpdateMessage, del chan DeleteMessage) {
	for {
		select {
		case <-ctx.Done():
			db.log.Info("UpdateListener context canceled")
			return
		case updMsg := <-upd:
			db.log.Info("Got messages:%d, users:%d, chats:%d updated", len(updMsg.Message), len(updMsg.User), len(updMsg.Chat))
		case delMsg := <-del:
			db.log.Info("Got messages:%d, users:%d, chats:%d deleted", len(delMsg.Message), len(delMsg.User), len(delMsg.Chat))
			return
		}
	}
}
