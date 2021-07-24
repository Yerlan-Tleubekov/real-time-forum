package repository

import (
	"database/sql"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
	"github.com/gorilla/websocket"
)

type IClient interface {
	NewClient(int, *Hub, *websocket.Conn, chan []byte) *Client
}

type DialogRoom interface {
	GetAllRooms(*models.UsedID) ([]*models.DialogRoom, error)
	GetDialogRoomByID(int) (*models.DialogRoom, error)
	CreateDialogRoom(int, int) error
	GetChatHubs() *ChatHubs
}
type UserSessionsActivity interface {
	SaveToken(int, string) error
	DeleteToken(int) error
	GetToken(int) (string, error)
}

type User interface {
	CreateUser(*models.User) (int64, error)
	GetUserByEmail(string) (*models.User, error)
	GetUserByNickname(string) (*models.User, error)
	GetUserByID(int) (*models.User, error)
	UserSessionsActivity
}

type Comment interface {
	CreateComment(*models.Comment) (int64, error)
	DeleteComment(int) (int64, error)
}

type Message interface {
	CreateMessage(*models.MessageHistory) (int64, error)
}

type Repository struct {
	User
	Comment
	DialogRoom
	Message
	IClient
}

func NewRepository(db *sql.DB, sessionTokens *SessionTokens, chatHubs *ChatHubs) *Repository {
	return &Repository{
		User:       NewUserRepository(db, sessionTokens),
		Comment:    NewCommentRepository(db, sessionTokens),
		DialogRoom: NewDialogRoomRepository(db, sessionTokens, chatHubs),
		Message:    NewMessageRepository(db, sessionTokens),
		IClient:    NewClientRepository(db),
	}
}
