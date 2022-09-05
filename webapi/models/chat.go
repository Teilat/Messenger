package models

type Chat struct {
	Name        string  `json:"name"     example:"Super Chat"`
	CreatedAt   string  `json:"createdAt" example:"1662070156" description:"unix time"`
	LastMessage Message `json:"lastMessage"`
}
