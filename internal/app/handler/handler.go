package handler

import (
	"net/http"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/server"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {

	return &Handler{services}
}

func (h *Handler) InitHandler() *http.ServeMux {

	fs := http.FileServer(http.Dir("./web/dist"))

	contTypeJson := "application/json"

	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", fs))
	mux.Handle("/auth/sign-up", server.Middlewares(h.SignUp(), []string{http.MethodPost, http.MethodOptions}, contTypeJson, false))
	mux.Handle("/auth/sign-in", server.Middlewares(h.SignIn(), []string{http.MethodPost, http.MethodOptions}, contTypeJson, false))
	mux.HandleFunc("/auth/logout", h.Logout)
	mux.HandleFunc("/comment/create", h.CreateComment)
	mux.HandleFunc("/comment/delete", h.DeleteComment)
	mux.HandleFunc("/dialog-rooms/all-rooms", h.GetAllRooms)
	mux.HandleFunc("/dialog-rooms/create-room", h.CreateRoom)
	mux.HandleFunc("/dialog-rooms/chat", h.CreateMessage)

	return mux
}
