package cache

import (
	"Messenger/internal/database"
	"Messenger/internal/logger"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// test for CRUD
func Test_Chat(t *testing.T) {
	c, upd, del := NewCache(logger.NewLogger("[Test Cache]"),
		func() ([]*database.User, []*database.Message, []*database.Chat) {
			return []*database.User{
				{Id: uuid.MustParse("3c54480e-d09a-11ed-afa1-0242ac120002"), Name: "new user 1"},
				{Id: uuid.MustParse("3c544cbe-d09a-11ed-afa1-0242ac120002"), Name: "new user 2"},
			}, []*database.Message{}, []*database.Chat{}
		}).Start()
	database.NewDbProvider(logger.NewLogger("[Test Database]")).StartUpdateListener(context.TODO(), upd, del)

	chat := &database.Chat{
		Id:   uuid.New(),
		Name: "Test chat",
		Users: []*database.User{
			{Id: uuid.MustParse("3c54480e-d09a-11ed-afa1-0242ac120002"), Name: "new user 1"},
			{Id: uuid.MustParse("3c544cbe-d09a-11ed-afa1-0242ac120002"), Name: "new user 2"},
		},
		Messages: []database.Message{},
	}
	// create chat
	err := c.CreateChat(chat)
	assert.NoErrorf(t, err, "CreateChat error:%v", err)

	createdChat, ok := c.Chat(chat.Id)
	if !ok {
		t.Error("Get chat error")
	}
	assert.Equalf(t, chat, createdChat, "Chats not equal \nExpected:%+v, \nGot:%+v", chat, createdChat)

	// update chat
	chat.Name = "Updated test chat"
	err = c.UpdateChat(chat)
	assert.NoErrorf(t, err, "UpdateChat error:%v", err)

	updatedChat, ok := c.Chat(chat.Id)
	if !ok {
		t.Error("Get chat error")
	}
	assert.Equalf(t, chat, updatedChat, "Chats not equal \nExpected:%+v, \nGot:%+v", chat, updatedChat)

	// delete chat
	err = c.DeleteChat(chat.Id)
	assert.NoErrorf(t, err, "CreateChat error:%v", err)

	deletedChat, ok := c.Chat(chat.Id)
	if ok {
		t.Error("Get deleted chat error")
	}
	assert.Equalf(t, (*database.Chat)(nil), deletedChat, "Chats not equal \nExpected:%+v, \nGot:%+v", nil, deletedChat)
}

func Test_cache_ChatsByUser(t *testing.T) {
	tests := []struct {
		name   string
		userId uuid.UUID
		want   []uuid.UUID
	}{
		{
			name:   "user 1",
			userId: uuid.MustParse("195d2788-e675-4d62-b3b5-ba6a75261485"),
			want:   []uuid.UUID{uuid.MustParse("c91d7e4e-a33f-4dd3-91ef-bb8fa9904de5")}},
		{
			name:   "user 2",
			userId: uuid.MustParse("ba96f5d5-f00b-436d-a79f-605dcef1ad31"),
			want:   []uuid.UUID{uuid.MustParse("db204477-f936-48a0-8ca6-d9fb67ac52f7")}},
		{
			name:   "user 3",
			userId: uuid.MustParse("bc563fb6-1fc3-44e7-832b-c4f7566f0c94"),
			want: []uuid.UUID{
				uuid.MustParse("db204477-f936-48a0-8ca6-d9fb67ac52f7"),
				uuid.MustParse("c91d7e4e-a33f-4dd3-91ef-bb8fa9904de5"),
			},
		},
	}
	c, _, _ := NewCache(logger.NewLogger("[Test Cache]"),
		func() ([]*database.User, []*database.Message, []*database.Chat) {
			return []*database.User{
					{Id: uuid.MustParse("195d2788-e675-4d62-b3b5-ba6a75261485"), Name: "User 1"},
					{Id: uuid.MustParse("ba96f5d5-f00b-436d-a79f-605dcef1ad31"), Name: "User 2"},
					{Id: uuid.MustParse("bc563fb6-1fc3-44e7-832b-c4f7566f0c94"), Name: "User 3"},
				},
				[]*database.Message{},
				[]*database.Chat{
					{Id: uuid.MustParse("c91d7e4e-a33f-4dd3-91ef-bb8fa9904de5"), Name: "chat 1",
						Users: []*database.User{
							{Id: uuid.MustParse("195d2788-e675-4d62-b3b5-ba6a75261485"), Name: "user 1"},
							{Id: uuid.MustParse("bc563fb6-1fc3-44e7-832b-c4f7566f0c94"), Name: "user 3"},
						},
					},
					{Id: uuid.MustParse("db204477-f936-48a0-8ca6-d9fb67ac52f7"), Name: "chat 2",
						Users: []*database.User{
							{Id: uuid.MustParse("ba96f5d5-f00b-436d-a79f-605dcef1ad31"), Name: "user 2"},
							{Id: uuid.MustParse("bc563fb6-1fc3-44e7-832b-c4f7566f0c94"), Name: "user 3"},
						},
					},
				}
		}).Start()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, _ := c.User(tt.userId)
			want := make([]*database.Chat, 0)
			for _, u := range tt.want {
				chat, _ := c.Chat(u)
				want = append(want, chat)
			}
			assert.Equalf(t, want, c.ChatsByUser(user), "ChatsByUser(%v)", user)
		})
	}
}

func Test_containsUser(t *testing.T) {
	tests := []struct {
		name  string
		users []*database.User
		user  *database.User
		want  bool
	}{
		{
			name: "true",
			users: []*database.User{
				{Id: uuid.MustParse("b862afec-d0b5-11ed-afa1-0242ac120002"), Name: "User 1"},
				{Id: uuid.MustParse("b862b38e-d0b5-11ed-afa1-0242ac120002"), Name: "User 2"},
				{Id: uuid.MustParse("b862b5d2-d0b5-11ed-afa1-0242ac120002"), Name: "User 3"}},
			user: &database.User{Id: uuid.MustParse("b862b38e-d0b5-11ed-afa1-0242ac120002"), Name: "User 2"},
			want: true,
		},
		{
			name: "false",
			users: []*database.User{
				{Id: uuid.MustParse("b862afec-d0b5-11ed-afa1-0242ac120002"), Name: "User 1"},
				{Id: uuid.MustParse("b862b38e-d0b5-11ed-afa1-0242ac120002"), Name: "User 2"},
				{Id: uuid.MustParse("b862b5d2-d0b5-11ed-afa1-0242ac120002"), Name: "User 3"}},
			user: &database.User{Id: uuid.MustParse("b862b48e-d0b5-11ed-afa1-0242ac120002"), Name: "User 4"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, containsUser(tt.users, tt.user), "containsUser(%v, %v)", tt.users, tt.user)
		})
	}
}

func Test_deleteChatFromSlice(t *testing.T) {
	tests := []struct {
		name  string
		slice []*database.Chat
		id    uuid.UUID
		want  []*database.Chat
	}{
		{
			name: "",
			slice: []*database.Chat{
				{Id: uuid.MustParse("c91d7e4e-a33f-4dd3-91ef-bb8fa9904de5")},
				{Id: uuid.MustParse("db204477-f936-48a0-8ca6-d9fb67ac52f7")},
			},
			id: uuid.MustParse("c91d7e4e-a33f-4dd3-91ef-bb8fa9904de5"),
			want: []*database.Chat{
				{Id: uuid.MustParse("db204477-f936-48a0-8ca6-d9fb67ac52f7")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, deleteChatFromSlice(tt.slice, tt.id), "deleteChatFromSlice(%v, %v)", tt.slice, tt.id)
		})
	}
}
