package models

import "go-boilerplate/modules/core"

type User struct {
	core.Base `gorm:"embedded"`
	Username  string `json:"username" gorm:"size:255;not null;default:john_doe"`
	Password  string `json:"-" gorm:"size:255;not null"`
	Salt      string `json:"-" gorm:"size:255;not null"`
	Name      string `json:"name" gorm:"size:255;not null"`
}
