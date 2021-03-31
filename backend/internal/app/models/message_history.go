package models

import (
	"time"
)

type MessageHistory struct {
	ID           int       `json:"ID"`
	DialogRoomID int       `json:"dialogRoomID"`
	Message      string    `json:"message"`
	UserID       int       `json:"userID"`
	CreatedDate  time.Time `json:"createdDate"`
}
