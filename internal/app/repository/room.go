package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
)

type DialogRoomRepository struct {
	db            *sql.DB
	sessionTokens *SessionTokens
	Hubs          *ChatHubs
}

func NewDialogRoomRepository(db *sql.DB, sessionTokens *SessionTokens, chatHubs *ChatHubs) *DialogRoomRepository {
	return &DialogRoomRepository{db, sessionTokens, chatHubs}
}

func (dialogRoomRepo *DialogRoomRepository) GetAllRooms(userID *models.UsedID) ([]*models.DialogRoom, error) {
	dialogRooms := []*models.DialogRoom{}

	rows, err := dialogRoomRepo.db.Query(`
					SELECT * 
					FROM dialog_room 
					WHERE first_user_id = $1
					OR second_user_id = $1
	`, userID.UserID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		dialogRoom := new(models.DialogRoom)

		if err := rows.Scan(
			&dialogRoom.ID,
			&dialogRoom.FirstUserID,
			&dialogRoom.FirstUserName,
			&dialogRoom.SecondUserID,
			&dialogRoom.SecondUserName,
			&dialogRoom.CreatedDate,
			&dialogRoom.LastMessageDate); err != nil {
			return nil, err
		}
		dialogRooms = append(dialogRooms, dialogRoom)

	}

	return dialogRooms, nil
}

func (dialogRoomRepo *DialogRoomRepository) GetDialogRoomByID(dialogRoomID int) (*models.DialogRoom, error) {
	var dialogRoom models.DialogRoom

	row := dialogRoomRepo.db.QueryRow(`
				SELECT *
				FROM dialog_room
				WHERE id = $1
	`, dialogRoomID)

	err := row.Err()
	if err != nil {
		return nil, errors.New("Room has empty")
	}

	row.Scan(
		&dialogRoom.ID,
		&dialogRoom.FirstUserID,
		&dialogRoom.FirstUserName,
		&dialogRoom.SecondUserID,
		&dialogRoom.SecondUserName,
		&dialogRoom.CreatedDate,
		&dialogRoom.LastMessageDate,
	)

	return &dialogRoom, nil
}

func (dialogRoomRepo *DialogRoomRepository) CreateDialogRoom(firstUserID, secondUserID int) error {
	_, err := dialogRoomRepo.db.Exec(`
		INSERT INTO dialog_room (first_user_id, first_user_name, second_user_id, second_user_name, created_date, last_message_date)
		VALUES (
							$1,
							(SELECT nickname FROM user WHERE id = $1),
							$2,
							(SELECT nickname FROM user WHERE id = $2),
							$3,
							$3)
	`, firstUserID, secondUserID, time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (dialogRoomRepo *DialogRoomRepository) GetChatHubs() *ChatHubs {
	return dialogRoomRepo.Hubs
}
