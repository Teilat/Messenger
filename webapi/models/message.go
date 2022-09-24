package models

type Message struct {
	Text      string
	CreatedAt string
	EditedAt  string
	User      string
}

type AddMessage struct {
	Text   string
	ChatId uint32
	User   string
}
