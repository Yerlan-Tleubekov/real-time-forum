package service

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
	"github.com/mattn/go-sqlite3"
)

func (commentService *UserService) CreateComment(comment *models.Comment) (error, int) {

	comment.Text = strings.Trim(comment.Text, " ")
	if comment.Text == "" {
		return errors.New("Comment not be empty"), http.StatusBadRequest
	}

	_, err := commentService.repo.Comment.CreateComment(comment)
	if _, ok := err.(sqlite3.Error); ok {

		return errors.New("error on create comment"), http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (service *UserService) DeleteComment(commentID int) (error, int) {

	return nil, http.StatusOK
}
