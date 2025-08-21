package models

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/users/models"

	"github.com/google/uuid"
)

type Message struct {
	core.Base `gorm:"embedded"`
	Content   string      `json:"content" gorm:"size:255"`
	SenderID  uuid.UUID   `json:"sender_id" gorm:"type:uuid;not null"`
	RoomID    uuid.UUID   `json:"room_id" gorm:"type:uuid;not null"`
	Sender    models.User `json:"sender" gorm:"foreignKey:SenderID"`
	Room      Room        `json:"room" gorm:"foreignKey:RoomID"`
}
