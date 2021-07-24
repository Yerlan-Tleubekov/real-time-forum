package service

import (
	"net/http"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
)

func (userS *UserService) GetUserByID(userID int) (*models.User, error, int) {
	var user *models.User

	user, err := userS.repo.User.GetUserByID(userID)

	if err != nil {
		return nil, err, http.StatusBadRequest
	}

	return user, nil, http.StatusOK

}
