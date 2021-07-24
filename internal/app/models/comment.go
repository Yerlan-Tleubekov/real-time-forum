package models

import "time"

type Comment struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userID"`
	PostID      int       `json:"postID"`
	Text        string    `json:"text"`
	CreatedDate time.Time `json:"createdDate"`
}
