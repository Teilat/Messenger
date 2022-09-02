package models

type Chat struct {
	Name            string `json:"name"     example:"Super Chat"`
	CreatedAt       string `json:"createdAt" example:"1662070156" description:"unix time"`
	LastMessage     string `json:"lastMessage" example:"last message" `
	LastMessageUser string `json:"lastMessageUser" example:"user" `
	LastMessageTime string `json:"lastMessageTime" example:"1662070156" description:"unix time"`
}
