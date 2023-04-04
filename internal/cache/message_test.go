package cache

import (
	"Messenger/internal/database"
	"Messenger/internal/logger"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Message(t *testing.T) {
	c, upd, del := NewCache(logger.NewLogger("[Test Cache]"),
		func() ([]*database.User, []*database.Message, []*database.Chat) {
			return []*database.User{
					{Id: uuid.MustParse("3c54480e-d09a-11ed-afa1-0242ac120002"), Name: "new user 1"},
					{Id: uuid.MustParse("3c544cbe-d09a-11ed-afa1-0242ac120002"), Name: "new user 2"},
				},
				[]*database.Message{},
				[]*database.Chat{{
					Id:   uuid.MustParse("c91d7e4e-a33f-4dd3-91ef-bb8fa9904de5"),
					Name: "chat 1",
					Users: []*database.User{
						{Id: uuid.MustParse("3c54480e-d09a-11ed-afa1-0242ac120002"), Name: "new user 1", Username: "user1"},
						{Id: uuid.MustParse("3c544cbe-d09a-11ed-afa1-0242ac120002"), Name: "new user 2", Username: "user2"},
					},
				}}
		}).Start()
	database.NewDbProvider(logger.NewLogger("[Test Database]")).StartUpdateListener(context.TODO(), upd, del)

	msg := &database.Message{
		Id:        uuid.New(),
		IntId:     0,
		Text:      "message text",
		CreatedAt: time.Now(),
		ChatId:    uuid.MustParse("c91d7e4e-a33f-4dd3-91ef-bb8fa9904de5"),
		Username:  "user1",
	}

	// create message
	err := c.CreateMessage(msg)
	assert.NoErrorf(t, err, "CreateMessage error:%v", err)

	createdMsg, ok := c.Message(msg.Id)
	if !ok {
		t.Error("Get message error")
	}
	assert.Equalf(t, msg, createdMsg, "Messages not equal \nExpected:%+v, \nGot:%+v", msg, createdMsg)

	// update message
	msg.Text = "updated text"
	msg.EditedAt = time.Now()
	err = c.UpdateMessage(msg)
	assert.NoErrorf(t, err, "UpdateMessage error:%v", err)

	updatedMessage, ok := c.Message(msg.Id)
	if !ok {
		t.Error("Get message error")
	}
	assert.Equalf(t, msg, updatedMessage, "Messages not equal \nExpected:%+v, \nGot:%+v", msg, createdMsg)

	// delete message
	msg.DeletedAt = time.Now()
	msg.DeletedForAll = true
	err = c.DeleteMessage(msg.Id)
	assert.NoErrorf(t, err, "DeleteMessage error:%v", err)

	deletedMsg, ok := c.Message(msg.Id)
	if ok {
		t.Error("Get deleted message error")
	}
	assert.Equalf(t, (*database.Message)(nil), deletedMsg, "Messages not equal \nExpected:%+v, \nGot:%+v", nil, deletedMsg)
}
