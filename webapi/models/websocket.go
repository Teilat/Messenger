package models

type MessageType struct {
	Action  string                 `json:"action"`
	Payload map[string]interface{} `json:"payload"`
}

type SendMessage struct {
	Text string `json:"text"`
}

type EditMessage struct {
	MessageId uint32 `json:"messageId"`
	NewText   string `json:"text"`
}

type DeleteMessage struct {
	MessageId uint32 `json:"messageId"`
}

type ReplyMessage struct {
	ReplyMessageId uint32      `json:"replyMessageId"`
	Message        SendMessage `json:"message"`
}
type GetMessages struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
