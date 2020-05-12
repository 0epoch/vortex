package model

import "time"

type User struct {
	ID int64 `json:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Avatar string `json:"avatar" form:"avatar"`
	LastAt time.Time `json:"last_at"`
	LastIP string `json:"last_ip" form:"last_ip"`
	CreatedIP string `json:"created_ip" form:"created_ip"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}