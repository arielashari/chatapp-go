package dto

import "github.com/google/uuid"

type CreateProfileDTO struct {
	Bio          *string   `json:"bio" binding:"omitempty"`
	ProfileImage *string   `json:"profile_image" binding:"omitempty"`
	UserID       uuid.UUID `json:"user_id" binding:"required" type:"uuid"`
}
