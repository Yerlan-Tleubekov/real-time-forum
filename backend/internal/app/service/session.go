package service

import (
	"errors"
	"net/http"
)

func (userS *UserService) GetToken(userID int) (string, error, int) {
	session, err := userS.repo.GetToken(userID)

	if err != nil {
		return "", err, http.StatusUnauthorized
	}

	return session, nil, http.StatusOK

}

func (userS *UserService) DeleteToken(userID int) (error, int) {
	return nil, http.StatusOK
}

func (userS *UserService) CompareSessions(requestedSession, savedSession string) (error, int) {
	if requestedSession != savedSession {
		return errors.New("Unauthorized"), http.StatusUnauthorized
	}

	return nil, http.StatusOK
}
