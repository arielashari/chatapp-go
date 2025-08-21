package models

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/users/models"

	"github.com/google/uuid"
)

type RoomMember struct {
	core.Base `gorm:"embedded"`
	RoomID    uuid.UUID   `gorm:"type:uuid;not null"`
	UserID    uuid.UUID   `gorm:"type:uuid;not null"`
	Room      Room        `gorm:"foreignKey:RoomID"`
	User      models.User `gorm:"foreignKey:UserID"`
}
