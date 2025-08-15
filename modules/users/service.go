package users

import (
	"crypto/sha512"
	"encoding/hex"
	"go-boilerplate/modules/core"
	ProfileService "go-boilerplate/modules/profile"
	ProfileDto "go-boilerplate/modules/profile/dto"
	"go-boilerplate/modules/users/dto"
	"go-boilerplate/modules/users/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/pbkdf2"
)

func GetAllUser(page string, pageSize string) (error, int64, []models.User) {
	limit, offset, err := core.Paging(page, pageSize)

	var count int64 = 0
	data := make([]models.User, 0)

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

func GetUserDetail(id string) (error, models.User) {
	var user models.User

	if result := core.DB.Where("id = ?", id).Take(&user); result.Error != nil {
		return result.Error, user
	}

	return nil, user
}

func CreateUser(createUserDto dto.UserDTO) (error, models.User) {
	var user models.User

	user.Name = createUserDto.Name
	user.Username = createUserDto.Username
	user.Salt = uuid.New().String()

	passwordHashed := pbkdf2.Key([]byte(createUserDto.Password), []byte(user.Salt), 4096, 100, sha512.New)

	user.Password = hex.EncodeToString(passwordHashed)

	if result := core.DB.Create(&user); result.Error != nil {
		return result.Error, user
	}

	profileDto := ProfileDto.CreateProfileDTO{
		UserID:       user.ID,
		ProfileImage: nil,
		Bio:          nil,
	}
	err, _ := ProfileService.CreateProfile(profileDto)

	if err != nil {
		return err, user
	}

	return nil, user
}

func UpdateUser(id string, updateUserDto dto.UserDTO) (error, models.User) {
	err, user := GetUserDetail(id)

	if err != nil {
		return err, user
	}

	user.Name = updateUserDto.Name

	if result := core.DB.Save(&user); result.Error != nil {
		return result.Error, user
	}

	return nil, user
}

func DeleteUser(id string) (error, models.User) {
	err, user := GetUserDetail(id)

	if err != nil {
		return err, user
	}

	if result := core.DB.Delete(&user); result.Error != nil {
		return result.Error, user
	}

	return nil, user
}
