package models

// Login is a simple type to contain username and password
type Login struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type Error struct {
	Content string `json:"content"`
}
