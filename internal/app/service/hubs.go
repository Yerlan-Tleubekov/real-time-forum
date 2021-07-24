package service

import (
	"errors"
	"log"
	"net/http"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/repository"
)

func (userS *UserService) Register(dialogRoomID int, hub *repository.Hub) (error, int) {
	userS.repo.DialogRoom.GetChatHubs().Register(dialogRoomID, hub)

	return nil, http.StatusOK
}

func (userS *UserService) GetHub(dialogRoomID int) (*repository.Hub, error, int) {
	hub, ok := userS.repo.DialogRoom.GetChatHubs().GetHub(dialogRoomID)
	if !ok {
		return nil, errors.New("empty"), http.StatusBadRequest
	}
	return hub, nil, http.StatusOK
}

func (userS *UserService) DeleteHub(dialogRoomID int) (error, int) {

	userS.repo.DialogRoom.GetChatHubs().Delete(dialogRoomID)

	return nil, http.StatusOK
}

func (userS *UserService) NewHub() *repository.Hub {
	return repository.NewHub()
}

func (userService *UserService) ServeWs(w http.ResponseWriter, r *http.Request, hub *repository.Hub, userID, dialogRoomID int) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := userService.repo.IClient.NewClient(userID, hub, conn, make(chan []byte, 256))

	// client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump(userID, dialogRoomID)
}
