package handler

import (
	"net/http"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {

	return &Handler{services}
}

func (h *Handler) InitHandler() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!\n"))
	})

	mux.HandleFunc("/auth/sign-up", h.SignUp)
	mux.HandleFunc("/auth/sign-in", h.SignIn)
	mux.HandleFunc("/auth/logout", h.Logout)
	mux.HandleFunc("/comment/create", h.CreateComment)
	mux.HandleFunc("/comment/delete", h.DeleteComment)
	mux.HandleFunc("/dialog-rooms/all-rooms", h.GetAllRooms)
	mux.HandleFunc("/dialog-rooms/create-room", h.CreateRoom)
	mux.HandleFunc("/dialog-rooms/chat", h.CreateMessage)

	return mux
}
