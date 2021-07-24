package models

import (
	"time"
)

// Client is a middleman between the websocket connection and the hub.

type DialogRoomUsers struct {
	UserID       int `json:"user_id"`
	SecondUserID int `json:"second_user_id"`
	DialogRoomID int `json:"dialog_room_id"`
}
type UsedID struct {
	UserID int `json:"user_id"`
}

type DialogRoomToUser struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	UserName        string    `json:"user_name"`
	CreatedDate     time.Time `json:"created_date"`
	LastMessageDate time.Time `json:"last_message_date"`
}

type DialogRoom struct {
	ID              int       `json:"id"`
	FirstUserID     int       `json:"first_user_id"`
	FirstUserName   string    `json:"first_user_name"`
	SecondUserID    int       `json:"second_user_id"`
	SecondUserName  string    `json:"second_user_name"`
	CreatedDate     time.Time `json:"created_date"`
	LastMessageDate time.Time `json:"last_message_date"`
}
