package model

import "time"

type UserGroup struct {
	UserID int64 `json:"user_id" form:"user_id"`
	GroupID int64 `json:"group_id" form:"group_id"`
	CreatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}