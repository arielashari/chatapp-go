package profile

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/profile/dto"
	"go-boilerplate/modules/profile/models"
)

func GetAllProfile(page string, pageSize string) (error, int64, []models.Profile) {
	limit, offset, err := core.Paging(page, pageSize)

	var count int64 = 0
	data := make([]models.Profile, 0)

	if err != nil {
		return err, count, data
	}

	if result := core.DB.Model(&data).Count(&count); result.Error != nil {
		return err, count, data
	}

	if result := core.DB.Limit(limit).Offset(offset).Order("created_at asc").Find(&data); result.Error != nil {
		return err, count, data
	}

	return nil, count, data
}

func GetProfileDetail(userId string) (error, models.Profile) {
	var profile models.Profile

	if result := core.DB.Where("user_id = ?", userId).Take(&profile); result.Error != nil {
		return result.Error, profile
	}

	return nil, profile
}

func CreateProfile(createProfileDto dto.CreateProfileDTO) (error, models.Profile) {
	var profile models.Profile

	if createProfileDto.Bio != nil {
		profile.Bio = createProfileDto.Bio
	}
	if createProfileDto.ProfileImage != nil {
		profile.ProfileImage = createProfileDto.ProfileImage
	}

	profile.UserID = createProfileDto.UserID

	if result := core.DB.Create(&profile); result.Error != nil {
		return result.Error, profile
	}

	return nil, profile
}

func UpdateProfile(id string, updateProfileDto dto.UpdateProfileDTO) (error, models.Profile) {
	err, profile := GetProfileDetail(id)

	if err != nil {
		return err, profile
	}

	//profile.Bio = updateProfileDto.Bio

	if result := core.DB.Save(&profile); result.Error != nil {
		return result.Error, profile
	}

	return nil, profile
}

func DeleteProfile(id string) (error, models.Profile) {
	err, profile := GetProfileDetail(id)

	if err != nil {
		return err, profile
	}

	if result := core.DB.Delete(&profile); result.Error != nil {
		return result.Error, profile
	}

	return nil, profile
}
