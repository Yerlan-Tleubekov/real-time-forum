package repository

import (
	"database/sql"
	"time"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
)

type CommentRepository struct {
	db            *sql.DB
	sessionTokens *SessionTokens
}

func NewCommentRepository(db *sql.DB, sessionTokens *SessionTokens) *CommentRepository {
	return &CommentRepository{db, sessionTokens}
}

func (commentRepo *CommentRepository) CreateComment(comment *models.Comment) (int64, error) {
	result, err := commentRepo.db.Exec(`
		INSERT INTO comment (user_id, post_id, text, created_date) values (?, ?, ?, ?)
	`, comment.UserID, comment.PostID, comment.Text, time.Now())
	if err != nil {
		return int64(0), err
	}

	return result.LastInsertId()
}

func (commentRepo *CommentRepository) DeleteComment(commentID int) (int64, error) {

	result, err := commentRepo.db.Exec(`DELETE FROM comment WHERE id = ?`, commentID)

	if err != nil {
		return int64(0), err
	}

	return result.RowsAffected()
}
