package handler

import (
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
)

type SignUpSuccess struct {
	UserID int `json:"id"`
}

type DialogRoom struct {
	DialogRooms []*models.DialogRoomToUser `json:"dialog_rooms"`
}
