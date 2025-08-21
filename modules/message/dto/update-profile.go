package dto

type UpdateProfileDTO struct {
	Bio          string `json:"bio" binding:"optional"`
	ProfileImage string `json:"profile_image" binding:"optional"`
	UserID       string `json:"user_id" binding:"required" type:"uuid"`
}
