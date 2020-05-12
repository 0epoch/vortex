package model

import "time"

const (
	//好友消息
	MSG_TYPE_F = 1
	//群消息
	MSG_TYPE_G = 2
)

type Message struct {
	ID int64 `json:"-"`
	FromUid int64 `json:"from_uid"`
	FromUsername string `json:"from_username"`
	ToID int64 `json:"to_id"`
	Content string `json:"content"`
	Type int `json:"type"`
	CreatedAt time.Time `json:"-"`
}