package service

import (
	"net/http"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/repository"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Session interface {
	GetToken(int) (string, error, int)
	DeleteToken(int) (error, int)
	CompareSessions(string, string) (error, int)
}
type Authorization interface {
	SignUp(*models.User) (error, int)
	SignIn(*models.UserSignIn) (*models.Session, error, int)
	ComparePasswords(*models.UserPasswords) error
	Logout(int) (error, int)
}

type Comment interface {
	CreateComment(*models.Comment) (error, int)
	DeleteComment(int) (error, int)
}

type DialogRoom interface {
	GetAllRooms(*models.UsedID) ([]*models.DialogRoomToUser, error, int)
	HasDialogRoomEmpty(int) (error, int)
	HasUserInDialogRoom(int, ...int) (error, int)
	GetDialogRoom(int) (*models.DialogRoom, error, int)
	CreateDialogRoom(int, int) (error, int)
	// ServeWs(*Hub, http.ResponseWriter, *http.Request)
	// NewHub() *Hub
}

type Message interface {
	CreateMessage(*models.MessageHistory) (error, int)
}

type ChatHubs interface {
	Register(int, *repository.Hub) (error, int)
	GetHub(int) (*repository.Hub, error, int)
	DeleteHub(int) (error, int)
	NewHub() *repository.Hub
	ServeWs(http.ResponseWriter, *http.Request, *repository.Hub, int, int)
}

type User interface {
	GetUserByID(int) (*models.User, error, int)
	Authorization
	Comment
	DialogRoom
	Message
	ChatHubs
	Session
}

type Service struct {
	User
}

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo}
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos),
	}
}
