package models

import (
	"database/sql/driver"
	"fmt"
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/users/models"

	"github.com/google/uuid"
)

type FriendRequestStatus string

const (
	Pending  FriendRequestStatus = "pending"
	Accepted FriendRequestStatus = "accepted"
	Rejected FriendRequestStatus = "rejected"
)

func (s *FriendRequestStatus) Scan(value interface{}) error {
	if value == nil {
		*s = Pending
		return nil
	}
	switch v := value.(type) {
	case string:
		*s = FriendRequestStatus(v)
	case []byte:
		*s = FriendRequestStatus(string(v))
	default:
		return fmt.Errorf("cannot scan FriendRequestStatus from %T", value)
	}
	return nil
}

func (s FriendRequestStatus) Value() (driver.Value, error) {
	return string(s), nil
}

type FriendRequests struct {
	core.Base
	SenderID   uuid.UUID           `json:"sender_id" gorm:"type:uuid;not null;column:sender_id"`
	ReceiverID uuid.UUID           `json:"receiver_id" gorm:"type:uuid;not null;column:receiver_id"`
	Status     FriendRequestStatus `json:"status" gorm:"type:friend_request_status;default:'pending'"`
	Sender     models.User         `json:"sender" gorm:"foreignKey:SenderID"`
	Receiver   models.User         `json:"receiver" gorm:"foreignKey:ReceiverID"`
}
