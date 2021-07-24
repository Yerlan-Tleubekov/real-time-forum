package models

import (
	"encoding/json"
	"errors"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/pkg/email"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/pkg/login"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/pkg/password"
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
	ID             int       `json:"id" db:"id"`
	Nickname       *string   `json:"nickname" db:"nickname"`
	Age            *int      `json:"age" db:"age"`
	Gender         *string   `json:"gender" db:"gender"`
	FirstName      *string   `json:"firstName" db:"firstName"`
	LastName       *string   `json:"lastName" db:"lastName"`
	Email          *string   `json:"email" db:"email"`
	Password       *string   `json:"password" db:"password"`
	PasswordRepeat *string   `json:"passwordRepeat,omitempty"`
	CreatedDate    time.Time `json:"createdDate,omitempty" db:"createdDate"`
	Role           *string   `json:"role,omitempty" db:"role"`
}

func (u *User) UnmarshalJSON(data []byte) error {
	type user User
	var a user
	var err error

	if err = json.Unmarshal(data, &a); err != nil {
		return err
	}

	*u = User(a)

	if err = u.Validate(); err != nil {
		return err
	}

	return nil
}

func (u *User) Validate() error {
	f := errors.New

	if u.Nickname == nil {
		return f("nickname is empty")
	}

	if *u.Nickname == "" || len(*u.Nickname) < 3 {
		return f("nickname is empty or len < 3")
	}

	if isValidLogin := login.CheckLogin(*u.Nickname); !isValidLogin {
		return f("incorrect login, only A-Za-z0-9(-,_)")
	}

	if u.Age == nil {
		return f("age is nil")
	}

	if *u.Age < 0 || *u.Age > 100 {
		return f("age < 0 || age > 100")
	}

	if u.Gender == nil {
		return f("gender is nil")
	}

	if *u.Gender == "" {
		return f("gender is empty")
	}

	if *u.Gender != "male" && *u.Gender != "female" {
		return f("invalid gender")
	}

	if u.FirstName == nil || u.LastName == nil {
		return f("firstname/lastname is nil")
	}

	if len(*u.FirstName) < 3 || len(*u.LastName) < 3 {
		return f("firstname/lastname len is smaller than 3")
	}

	if u.Email == nil {
		return f("email is nil")
	}

	if *u.Email == "" || !email.CheckValidEmail(*u.Email) {
		return f("incorrect email")
	}

	if u.Password == nil || u.PasswordRepeat == nil {
		return f("password is empty")
	}

	if password.CheckPasswordLength(*u.Password) || password.CheckPasswordLength(*u.PasswordRepeat) {
		return f("len password < 8")
	}

	return nil
}
