package model

import "time"

type UserFriend struct {
	UserID int64 `json:"user_id"`
	FriendID int64 `json:"friend_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}