package users

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/users/dto"
	"go-boilerplate/modules/users/models"
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

	if result := core.DB.Create(&user); result.Error != nil {
		return result.Error, user
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
