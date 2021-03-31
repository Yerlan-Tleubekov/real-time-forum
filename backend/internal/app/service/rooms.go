package service

import (
	"errors"
	"net/http"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
	sqlite "github.com/mattn/go-sqlite3"
)

func ternary(a, equal, b int) int {
	if a == equal {
		return b
	}
	return a
}

func ternaryString(a, equal int, firstName, secondName string) string {
	if a == equal {
		return secondName
	}
	return firstName
}

func (userService *UserService) GetAllRooms(userID *models.UsedID) ([]*models.DialogRoomToUser, error, int) {

	var dialogRooms []*models.DialogRoomToUser
	rooms, err := userService.repo.DialogRoom.GetAllRooms(userID)

	for i := 0; i < len(rooms); i++ {
		roomFromDB := rooms[i]

		room := &models.DialogRoomToUser{
			ID:              roomFromDB.ID,
			UserID:          ternary(roomFromDB.FirstUserID, userID.UserID, roomFromDB.SecondUserID),
			UserName:        ternaryString(roomFromDB.FirstUserID, userID.UserID, roomFromDB.FirstUserName, roomFromDB.SecondUserName),
			CreatedDate:     roomFromDB.CreatedDate,
			LastMessageDate: roomFromDB.LastMessageDate,
		}
		dialogRooms = append(dialogRooms, room)
	}

	if err != nil {
		return nil, err, http.StatusBadRequest
	}
	return dialogRooms, nil, http.StatusOK

}

func (userService *UserService) HasDialogRoomEmpty(dialogRoomID int) (error, int) {

	if _, err := userService.repo.DialogRoom.GetDialogRoomByID(dialogRoomID); err != nil {
		return err, http.StatusBadRequest
	}

	return nil, http.StatusOK
}

func (userService *UserService) HasUserInDialogRoom(dialogRoomID int, users ...int) (error, int) {
	dialogRoom, err := userService.repo.GetDialogRoomByID(dialogRoomID)

	if err != nil {
		return err, http.StatusBadRequest
	}

	for i := 0; i < len(users); i++ {

		if dialogRoom.FirstUserID != users[i] && dialogRoom.SecondUserID != users[i] {
			return errors.New("Incorrect user"), http.StatusBadRequest
		}

	}

	return nil, http.StatusOK
}

func (userService *UserService) GetDialogRoom(dialogRoomID int) (*models.DialogRoom, error, int) {
	dialogRoom, err := userService.repo.DialogRoom.GetDialogRoomByID(dialogRoomID)

	if err != nil {
		return nil, err, http.StatusBadRequest
	}

	return dialogRoom, nil, http.StatusOK
}

func (userS *UserService) CreateDialogRoom(firstUser, secondUser int) (error, int) {
	if sqliteError := userS.repo.DialogRoom.CreateDialogRoom(firstUser, secondUser); sqliteError != nil {
		if err, ok := sqliteError.(sqlite.Error); ok {
			if err.ExtendedCode == sqlite.ErrConstraintUnique {
				return errors.New("uje est' takoi room"), http.StatusBadRequest
			}
		}

		return nil, http.StatusOK
	}

	return nil, http.StatusOK
}
