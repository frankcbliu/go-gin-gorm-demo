package models

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model

	ID       int    `json:"id" gorm:"type:int"`
	Username string `json:"username" gorm:"type:varchar(30)"`
	Password string `json:"title" gorm:"type:varchar(100)"`
}

func (user *User) CreateUser(username string, password string) bool {
	result := SqliteDB.Create(&User{Username: username, Password: getMD5String(password)})
	if result.Error != nil {
		log.Println("SqliteDB.Create error", result.Error)
		return false
	}
	return true
}

func (user *User) CheckPassword(password string) bool {
	if len(password) <= 0 {
		log.Println("password empty")
		return false
	}
	if user.Password != getMD5String(password) {
		log.Println("password md5 check error")
		return false
	}
	return true
}

func (user *User) FindUser(username string) bool {
	SqliteDB.Find(&user, "username = ?", username)
	if user.ID > 0 {
		return true
	}
	log.Println("Not found user:", username)
	return false
}

// 获取 md5 计算值
func getMD5String(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}
