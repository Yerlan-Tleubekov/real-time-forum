package models

type PostCategories struct {
	ID         int `json:"id"`
	PostID     int `json:"postID"`
	CategoryID int `json:"categoryID"`
}
