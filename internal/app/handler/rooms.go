package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/pkg/response"
)

func (h *Handler) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	var answerJSON response.AnswerJSON
	var userID models.UsedID
	var dialogRooms DialogRoom
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, http.StatusBadRequest, answerJSON)
			return
		}

		if err := json.Unmarshal(body, &userID); err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, http.StatusBadRequest, answerJSON)
			return
		}

		rooms, err, code := h.services.GetAllRooms(&userID)
		if err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, code, answerJSON)
			return
		}

		dialogRooms = DialogRoom{DialogRooms: rooms}
		answerJSON = response.AnswerJSON{Data: dialogRooms, Code: code}
		response.WriteResponse(w, code, answerJSON)

		return
	default:
		errMessage := fmt.Sprintf("%s method not allowed", r.Method)
		errorJSON := &response.ErrorJson{Message: errMessage}
		answerJSON = response.AnswerJSON{Data: errorJSON, Code: http.StatusBadRequest}
		response.WriteResponse(w, http.StatusBadRequest, answerJSON)
		return

	}
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var userID models.DialogRoomUsers

	if r.Method != "POST" {
		err := fmt.Sprintf("%s method not allowed", r.Method)
		response.ErrorJsonWriter(errors.New(err), http.StatusMethodNotAllowed, w)

		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.ErrorJsonWriter(err, http.StatusBadRequest, w)
		return
	}

	if err := json.Unmarshal(body, &userID); err != nil {
		response.ErrorJsonWriter(err, http.StatusBadRequest, w)

		return
	}

	session, err := r.Cookie("forum")
	if err != nil {

		response.ErrorJsonWriter(err, http.StatusUnauthorized, w)
		return
	}

	userSavedSession, err, code := h.services.User.GetToken(userID.UserID)

	if err != nil {
		response.ErrorJsonWriter(err, code, w)
		return
	}

	if err, code := h.services.User.CompareSessions(session.Value, userSavedSession); err != nil {
		response.ErrorJsonWriter(err, code, w)

		return
	}

	if err, code := h.services.User.CreateDialogRoom(userID.UserID, userID.SecondUserID); err != nil {
		response.ErrorJsonWriter(err, code, w)
		return
	}

	response.WriteResponse(w, http.StatusOK, response.AnswerJSON{Data: "success", Code: http.StatusOK})

}
