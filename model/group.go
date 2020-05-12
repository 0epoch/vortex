package model

import "time"

type Group struct {
	ID int64 `json:"id"`
	Name string `json:"name" form:"name"`
	CreatedUid int64 `json:"created_uid" form:"created_uid"`
	MemberTotal int `json:"member_total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}