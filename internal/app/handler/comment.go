package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/pkg/response"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	var answerJSON response.AnswerJSON
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, http.StatusBadRequest, answerJSON)
			return
		}

		if err := json.Unmarshal(body, &comment); err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, http.StatusBadRequest, answerJSON)
			return
		}

		if err, code := h.services.CreateComment(&comment); err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, code, answerJSON)
			return
		}

		return
	default:

		return
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

}
