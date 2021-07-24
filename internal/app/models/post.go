package models

import "time"

type Post struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userID"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CreatedDate time.Time `json:"createdDate"`
}
