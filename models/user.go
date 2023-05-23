package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID       int    `gorm:"serializer:json"`
	Username string `gorm:"serializer:json"`
	Password string `gorm:"serializer:json"`
}

func (user *User) CreateUser(username string, password string) {
	SqliteDB.Create(&User{Username: username, Password: password})
}
