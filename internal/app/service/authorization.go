package service

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/pkg/database"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/pkg/password"
	sqlite "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

func (userService *UserService) SignUp(user *models.User) (error, int) {

	// start hashing password
	hashedPassword, err := password.HashPassword(*user.Password)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	*user.Password = hashedPassword

	_, err = userService.repo.CreateUser(user)
	if sqliteErr, ok := err.(sqlite.Error); ok {
		colName := database.DetectUniqueRowColName(sqliteErr.Error())
		currentError := errors.New("incorrect " + colName)

		if sqliteErr.ExtendedCode == sqlite.ErrConstraintUnique {
			return currentError, http.StatusBadRequest
		}
		return errors.New("user already exists"), http.StatusBadRequest

	}

	return nil, http.StatusOK
}

func (userService *UserService) SignIn(user *models.UserSignIn) (*models.Session, error, int) {
	var session *models.Session
	var userFromDB *models.User
	var err error
	emailDetector := "@"

	if strings.Contains(user.Login, emailDetector) {
		userFromDB, err = userService.repo.GetUserByEmail(user.Login)
	} else {
		userFromDB, err = userService.repo.GetUserByNickname(user.Login)
	}
	if err != nil {
		return nil, err, http.StatusBadRequest
	}

	if err = password.ComparePasswords(*userFromDB.Password, user.Password); err != nil {
		return nil, err, http.StatusBadRequest
	}

	session = &models.Session{
		UserID:  userFromDB.ID,
		ExpTime: time.Now().Add(time.Minute * 30),
		Token:   uuid.NewV4().String(),
	}

	// block for create session
	//
	userService.repo.SaveToken(userFromDB.ID, session.Token)
	//
	// end block for create session

	return session, nil, http.StatusOK

}

func (userService *UserService) Logout(userID int) (error, int) {
	if userID == 0 {
		return errors.New("user_id null"), http.StatusBadRequest
	}

	if err := userService.repo.DeleteToken(userID); err != nil {
		return err, http.StatusBadRequest
	}

	return nil, http.StatusOK
}

func (userService *UserService) ComparePasswords(user *models.User) error {
	isEqual := 0

	answer := strings.Compare(*user.Password, *user.PasswordRepeat)

	if answer == isEqual {
		return nil
	}

	err := errors.New("passwords is not equal")

	return err

}
