package models

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/users/models"

	"github.com/google/uuid"
)

type Room struct {
	core.Base `gorm:"embedded"`
	Name      *string     `json:"name" gorm:"size:255"`
	OwnerID   uuid.UUID   `json:"owner_id" gorm:"type:uuid;not null"`
	Owner     models.User `json:"owner" gorm:"foreignKey:OwnerID"`
}
