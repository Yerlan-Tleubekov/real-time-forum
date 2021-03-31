package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
)

type UserRepository struct {
	db            *sql.DB
	sessionTokens *SessionTokens
}

func NewUserRepository(db *sql.DB, sessionTokens *SessionTokens) *UserRepository {
	return &UserRepository{db, sessionTokens}
}

func (userRepos *UserRepository) CreateUser(user *models.User) (int64, error) {
	result, err := userRepos.db.Exec(`
	INSERT INTO user (nickname, age, gender, first_name, last_name, email,password, role,created_date) 
	VALUES (?,?,?,?,?, ? ,? ,? ,? )`,
		user.Nickname,
		user.Age,
		user.Gender,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password, "user", time.Now())

	if err != nil {
		return int64(0), err
	}
	return result.LastInsertId()

}

func (userRepos *UserRepository) GetUserByEmail(login string) (*models.User, error) {
	var user models.User

	rows := userRepos.db.QueryRow(`SELECT * FROM user WHERE email = ?`, login)

	if err := rows.Scan(
		&user.ID,
		&user.Nickname,
		&user.Age,
		&user.Gender,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedDate,
		&user.Role); err != nil {

		return nil, err
	}

	return &user, nil
}

func (userRepos *UserRepository) GetUserByNickname(login string) (*models.User, error) {
	var user models.User

	rows := userRepos.db.QueryRow(`SELECT * FROM user WHERE nickname = ?`, login)

	if err := rows.Scan(&user.ID,
		&user.Nickname,
		&user.Age,
		&user.Gender,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedDate,
		&user.Role); err != nil {

		return nil, err
	}

	return &user, nil
}

func (userRepos *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User

	rows := userRepos.db.QueryRow(`SELECT * FROM user WHERE id = ?`, id)

	if err := rows.Scan(&user.ID,
		&user.Nickname,
		&user.Age,
		&user.Gender,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedDate,
		&user.Role); err != nil {

		return nil, err
	}

	return &user, nil
}

func (userRepos *UserRepository) GetToken(userID int) (string, error) {
	token := userRepos.sessionTokens.sessionTokens[userID]

	if token == "" {
		return "", errors.New("Unauthorized")
	}

	return token, nil

}

func (userRepos *UserRepository) SaveToken(userID int, token string) error {
	userRepos.sessionTokens.sessionTokens[userID] = token
	fmt.Println(userRepos.sessionTokens)
	return nil
}
func (userRepos *UserRepository) DeleteToken(userID int) error {
	delete(userRepos.sessionTokens.sessionTokens, userID)
	fmt.Println(userRepos.sessionTokens)
	return nil
}
