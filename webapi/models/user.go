package models

type User struct {
	Username   string `json:"name"     example:"John"`
	Nickname   string `json:"nickname" example:"Nickname"`
	Bio        string `json:"bio"   example:"What are you taking about?" `
	Phone      string `json:"phone" example:"+78005553535"`
	CreatedAt  string `json:"createdAt" example:"1662070156" description:"unix time"`
	LastOnline string `json:"lastOnline" example:"1662070156" description:"unix time"`
	Chats      []Chat `json:"chats"`
}

type UserCredentials struct {
	Nickname string `json:"name" example:"Nickname"`
	Password string `json:"password" example:"password"`
}

type AddUser struct {
	Username string `json:"login" form:"login" example:"User"`
	Nickname string `json:"name" form:"name" example:"User"`
	Password string `json:"password" form:"password" example:"password"`
	Phone    string `json:"phone" form:"phone" example:"+78005553535"`
}

type UpdateUser struct {
	Name     string `json:"name"     example:"John"`
	Surname  string `json:"surname"  example:"Joe"`
	Nickname string `json:"nickname" example:"Nickname"`
}

type DeleteUser struct {
	Username string `json:"name" example:"Admin"`
}

type UpdateUserPassword struct {
	Id       int32  `json:"id" example:"15" format:"integer"`
	Password string `json:"password" example:"password"`
}
