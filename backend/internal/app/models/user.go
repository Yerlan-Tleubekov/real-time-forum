package models

import (
	"time"
)

type Session struct {
	UserID  int       `json:"user_id"`
	Token   string    `json:"token"`
	ExpTime time.Time `json:"exp_time"`
}

type UserSignIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserLogout struct {
	UserID int `json:"user_id"`
}

type UserPasswords struct {
	Password       string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
}

type User struct {
	ID          int       `json:"id"`
	Nickname    string    `json:"nickname"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedDate time.Time `json:"createdDate"`
	Role        string    `json:"role"`
}
