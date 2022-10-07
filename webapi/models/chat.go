package models

type Chat struct {
	Id        uint32    `json:"id"`
	Name      string    `json:"name"     example:"Super Chat"`
	CreatedAt string    `json:"createdAt" example:"1662070156" description:"unix time"`
	Messages  []Message `json:"messages"`
	Users     []string  `json:"users"`
}

type AddChat struct {
	Name  string   `json:"name"     example:"Super Chat"`
	Users []string `json:"users"`
}

type AddUserToChat struct {
	Name   string `json:"name"   example:"Super Chat"`
	UserId string `json:"userId" example:"Admin"`
}
