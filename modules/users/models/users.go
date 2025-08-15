package models

import "go-boilerplate/modules/core"

type User struct {
	core.Base `gorm:"embedded"`
	Email     string `json:"email" gorm:"size:255;not null"`
	Password  string `json:"-" gorm:"size:255;not null"`
	Salt      string `json:"-" gorm:"size:255;not null"`
	Name      string `json:"name" gorm:"size:255;not null"`
}
