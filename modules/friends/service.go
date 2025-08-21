package friends

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/friends/models"

	"github.com/google/uuid"
)

func GetAllFriends(userId string, page string, pageSize string) (error, int64, []models.Friends) {
	limit, offset, err := core.Paging(page, pageSize)

	var count int64
	data := make([]models.Friends, 0)

	if err != nil {
		return err, count, data
	}

	condition := core.DB.Where("user_id_1 = ? OR user_id_2 = ?", userId, userId).Preload("User1").Preload("User2")

	if result := condition.Model(&models.Friends{}).Count(&count); result.Error != nil {
		return result.Error, count, data
	}

	if result := condition.Limit(limit).Offset(offset).Order("created_at asc").Find(&data); result.Error != nil {
		return result.Error, count, data
	}

	return nil, count, data
}

func SendFriendRequest(userId string, friendId string) (error, models.FriendRequests) {
	var friendRequest models.FriendRequests

	friendRequest.SenderID, _ = uuid.Parse(userId)
	friendRequest.ReceiverID, _ = uuid.Parse(friendId)
	friendRequest.Status = "pending"
	if result := core.DB.Create(&friendRequest); result.Error != nil {
		return result.Error, friendRequest
	}

	return nil, friendRequest
}

func GetFriendRequests(userId string, page string, pageSize string) (error, int64, []models.FriendRequests) {
	limit, offset, err := core.Paging(page, pageSize)

	var count int64 = 0
	data := make([]models.FriendRequests, 0)

	if err != nil {
		return err, count, data
	}

	condition := core.DB.Where("receiver_id = ? AND status = 'pending'", userId).Preload("Sender").Preload("Receiver")

	if result := condition.Model(&models.FriendRequests{}).Count(&count); result.Error != nil {
		return result.Error, count, data
	}

	if result := condition.Limit(limit).Offset(offset).Order("created_at asc").Find(&data); result.Error != nil {
		return result.Error, count, data
	}

	return nil, count, data
}

func AcceptFriendRequest(userId string, friendId string) (error, models.FriendRequests) {
	var friendRequest models.FriendRequests

	if result := core.DB.Where("sender_id = ? AND receiver_id = ?", friendId, userId).First(&friendRequest); result.Error != nil {
		return result.Error, friendRequest
	}

	friendRequest.Status = "accepted"
	if result := core.DB.Save(&friendRequest); result.Error != nil {
		return result.Error, friendRequest
	}

	var friend models.Friends
	friend.UserID1, _ = uuid.Parse(userId)
	friend.UserID2, _ = uuid.Parse(friendId)
	if result := core.DB.Create(&friend); result.Error != nil {
		return result.Error, friendRequest
	}

	return nil, friendRequest
}

func RejectFriendRequest(userId string, friendId string) (error, models.FriendRequests) {
	var friendRequest models.FriendRequests

	if result := core.DB.Where("sender_id = ? AND receiver_id = ?", friendId, userId).First(&friendRequest); result.Error != nil {
		return result.Error, friendRequest
	}

	friendRequest.Status = "rejected"
	if result := core.DB.Save(&friendRequest); result.Error != nil {
		return result.Error, friendRequest
	}

	var friend models.Friends
	friend.UserID1, _ = uuid.Parse(userId)
	friend.UserID2, _ = uuid.Parse(friendId)
	if result := core.DB.Create(&friend); result.Error != nil {
		return result.Error, friendRequest
	}

	return nil, friendRequest
}

func DeleteFriend(userId string, friendId string) (error, models.Friends) {
	var friend models.Friends

	if result := core.DB.Where("user_id_1 = ? AND user_id_2 = ?", userId, friendId).First(&friend); result.Error != nil {
		return result.Error, friend
	}

	if result := core.DB.Delete(&friend); result.Error != nil {
		return result.Error, friend
	}
	return nil, friend
}
