package models

import (
	"go-boilerplate/modules/core"

	"github.com/google/uuid"
)

type Profile struct {
	core.Base    `gorm:"embedded"`
	ProfileImage *string   `json:"profile_image" gorm:"size:255"`
	Bio          *string   `json:"bio" gorm:"size:255"`
	UserID       uuid.UUID `json:"user_id" gorm:"type:uuid"`
}
