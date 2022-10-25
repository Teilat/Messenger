package cache

type Cache interface {
	CreateChat()
	UpdateChat()
	DeleteChat()

	CreateMessage()
	UpdateMessage()
	DeleteMessage()

	CreateUser()
	UpdateUser()
	DeleteUser()
}
