package resolver

import (
	"Messenger/database"
	"gorm.io/gorm"
	"time"
)

func GetUserByUsername(db *gorm.DB, username string) database.User {
	user := database.User{Username: username}
	db.Find(&user)
	updateLastOnline(db, username)
	return user
}

func updateLastOnline(db *gorm.DB, username string, tz ...time.Time) {
	updateTime := time.Now()
	if len(tz) > 0 {
		updateTime = tz[0]
	}
	db.Model(&database.User{Username: username}).Update("last_online", updateTime)
}
