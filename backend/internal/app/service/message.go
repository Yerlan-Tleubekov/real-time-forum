package service

import (
	"net/http"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
)

func (userService *UserService) CreateMessage(message *models.MessageHistory) (error, int) {

	return nil, http.StatusOK
}
