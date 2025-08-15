package models

import (
	"go-boilerplate/modules/core"

	"github.com/google/uuid"
)

type Friends struct {
	core.Base `gorm:"embedded"`
	UserID1   uuid.UUID `json:"user_id_1" gorm:"type:uuid"`
	UserID2   uuid.UUID `json:"user_id_2" gorm:"type:uuid"`
}
