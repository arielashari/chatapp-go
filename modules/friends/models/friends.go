package models

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/users/models"

	"github.com/google/uuid"
)

type Friends struct {
	core.Base `gorm:"embedded"`

	UserID1 uuid.UUID   `json:"user_id_1" gorm:"type:uuid;column:user_id_1;not null"`
	User1   models.User `json:"user_1" gorm:"foreignKey:UserID1"`
	UserID2 uuid.UUID   `json:"user_id_2" gorm:"type:uuid;column:user_id_2;not null"`
	User2   models.User `json:"user_2" gorm:"foreignKey:UserID2"`
}
