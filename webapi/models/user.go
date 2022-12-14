package models

type User struct {
	Username   string `json:"name"     example:"John"`
	Nickname   string `json:"nickname" example:"Nickname"`
	Bio        string `json:"bio"   example:"What are you taking about?" `
	Phone      string `json:"phone" example:"+78005553535"`
	CreatedAt  string `json:"createdAt" description:"str time"`
	LastOnline string `json:"lastOnline" description:"str time"`
	Image      []byte `json:"image" description:"user image"`
}

type AddUser struct {
	Username string `json:"login" form:"login" example:"User"`
	Nickname string `json:"name" form:"name" example:"User"`
	Password string `json:"password" form:"password" example:"password"`
	Phone    string `json:"phone" form:"phone" example:"+78005553535"`
}

type UpdateUser struct {
	Name     string `json:"name"     example:"John"`
	Nickname string `json:"nickname" example:"Nickname"`
	Bio      string `json:"bio" example:"Who are you talking about"`
}

type DeleteUser struct {
	Username string `json:"name" example:"Admin"`
}

type UpdateUserPassword struct {
	Id       int32  `json:"id" example:"15" format:"integer"`
	Password string `json:"password" example:"password"`
}
