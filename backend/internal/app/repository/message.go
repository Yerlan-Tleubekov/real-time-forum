package repository

import (
	"database/sql"
	"time"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
)

type MessageRepository struct {
	db            *sql.DB
	sessionTokens *SessionTokens
}

func NewMessageRepository(db *sql.DB, sessionToken *SessionTokens) *MessageRepository {
	return &MessageRepository{db, sessionToken}
}

func (messageRepo *MessageRepository) CreateMessage(message *models.MessageHistory) (int64, error) {

	result, err := messageRepo.db.Exec(`
			INSERT 
			FROM message_history (dialog_room_id, message, user_id, created_date) 
			VALUES($1, $2, $3, $4)
	`, message.DialogRoomID, message.Message, message.UserID, time.Now())

	if err != nil {
		return int64(0), err
	}

	return result.LastInsertId()
}
