package friends

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/friends/models"
)

func GetAllFriends(page string, pageSize string) (error, int64, []models.Friends) {
	limit, offset, err := core.Paging(page, pageSize)

	var count int64 = 0
	data := make([]models.Friends, 0)

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
