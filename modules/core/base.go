package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid"`
	CreatedAt time.Time       `json:"created_at" gorm:"default:now();autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"default:now();autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(_ *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
	return
}

func (base *Base) BeforeUpdate(_ *gorm.DB) (err error) {
	base.UpdatedAt = time.Now()
	return
}
