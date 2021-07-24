package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/repository"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/pkg/response"
)

func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var hub *repository.Hub
	var dialogRoomUsers *models.DialogRoomUsers

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// if r.Method != "POST" {
	// 	err := fmt.Sprintf("%s method not allowed", r.Method)
	// 	response.ErrorJsonWriter(errors.New(err), http.StatusBadRequest, w)
	// 	return
	// }

	// roomIDStr := r.URL.Query().Get("roomID")
	// userIDToStr := r.URL.Query().Get("userID")
	//
	// roomID, err := strconv.Atoi(roomIDStr)

	// if err != nil {
	// response.ErrorJsonWriter(err, http.StatusBadRequest, w)
	// return
	// }

	// userIdTo, err := strconv.Atoi(userIDToStr)

	// if err != nil {
	// response.ErrorJsonWriter(err, http.StatusBadRequest, w)
	// return
	// }

	// new
	userIDStr := r.URL.Query().Get("userID")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		response.ErrorJsonWriter(errors.New("user id"), http.StatusBadRequest, w)
		return
	}

	dialogRoomIDStr := r.URL.Query().Get("room")
	dialogRoomID, err := strconv.Atoi(dialogRoomIDStr)

	if err != nil {
		response.ErrorJsonWriter(errors.New("room id"), http.StatusBadRequest, w)
		return
	}

	secondUserIDStr := r.URL.Query().Get("secondUserID")
	secondUserID, err := strconv.Atoi(secondUserIDStr)

	dialogRoomUsers = &models.DialogRoomUsers{UserID: userID, SecondUserID: secondUserID, DialogRoomID: dialogRoomID}

	if err != nil {

		response.ErrorJsonWriter(errors.New("second user id"), http.StatusBadRequest, w)
		return
	}

	// session, err := r.Cookie("forum")
	// if err != nil {
	// 	fmt.Println("err in session")

	// 	response.ErrorJsonWriter(errors.New("Unauthorized"), http.StatusUnauthorized, w)
	// 	return
	// }

	// savedSession, err, code := h.services.GetToken(dialogRoomUsers.UserID)

	// if err != nil {
	// 	response.ErrorJsonWriter(err, code, w)
	// 	return
	// }

	// if err, code := h.services.User.CompareSessions(session.Value, savedSession); err != nil {
	// 	response.ErrorJsonWriter(err, code, w)

	// 	return
	// }

	if _, err, code := h.services.User.GetUserByID(dialogRoomUsers.SecondUserID); err != nil {
		response.ErrorJsonWriter(err, code, w)
		return
	}

	if err, code := h.services.User.HasDialogRoomEmpty(dialogRoomUsers.DialogRoomID); err != nil {
		response.ErrorJsonWriter(err, code, w)
		return
	}

	if err, code := h.services.User.HasUserInDialogRoom(dialogRoomUsers.DialogRoomID, dialogRoomUsers.UserID, dialogRoomUsers.SecondUserID); err != nil {
		response.ErrorJsonWriter(err, code, w)
		return
	}

	dialogRoom, err, code := h.services.User.GetDialogRoom(dialogRoomUsers.DialogRoomID)
	if err != nil {

		response.ErrorJsonWriter(err, code, w)
		return
	}

	hub, err, code = h.services.User.GetHub(dialogRoom.ID)

	if err != nil {
		hub = h.services.NewHub()
		h.services.User.Register(dialogRoom.ID, hub)
	}
	go hub.Run()

	h.services.User.ServeWs(w, r, hub, userID, dialogRoomID)

}
