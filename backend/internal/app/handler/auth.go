package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/pkg/response"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var answerJSON response.AnswerJSON
	var userPasswords models.UserPasswords
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, http.StatusBadRequest, answerJSON)
			return
		}

		if err := json.Unmarshal(body, &userPasswords); err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, http.StatusBadRequest, answerJSON)
			return
		}

		if err := h.services.User.ComparePasswords(&userPasswords); err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, http.StatusBadRequest, answerJSON)
			return
		}

		if err := json.Unmarshal(body, &user); err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, http.StatusBadRequest, answerJSON)

			return
		}

		if err, code := h.services.User.SignUp(&user); err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, code, answerJSON)

			return
		}

		successJSON := SignUpSuccess{UserID: user.ID}
		answerJSON = response.AnswerJSON{Data: successJSON, Code: http.StatusOK}
		response.WriteResponse(w, http.StatusOK, answerJSON)
		break
	default:
		errMessage := fmt.Sprintf("%s method not allowed", r.Method)
		errorJSON := &response.ErrorJson{Message: errMessage}
		answerJSON = response.AnswerJSON{Data: errorJSON, Code: http.StatusBadRequest}
		response.WriteResponse(w, http.StatusBadRequest, answerJSON)
		break

	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var answerJSON response.AnswerJSON
	var user models.UserSignIn
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, http.StatusBadRequest, answerJSON)
			return
		}

		if err := json.Unmarshal(body, &user); err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, http.StatusBadRequest, answerJSON)
			return
		}

		session, err, code := h.services.User.SignIn(&user)
		if err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, code, answerJSON)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "forum",
			Path:    "/",
			Value:   session.Token,
			Expires: session.ExpTime,
		})

		successJSON := SignUpSuccess{UserID: session.UserID}
		answerJSON = response.AnswerJSON{Data: successJSON, Code: http.StatusOK}
		response.WriteResponse(w, http.StatusOK, answerJSON)

	default:
		errMessage := fmt.Sprintf("%s method not allowed", r.Method)
		errorJSON := &response.ErrorJson{Message: errMessage}
		answerJSON = response.AnswerJSON{Data: errorJSON, Code: http.StatusBadRequest}
		response.WriteResponse(w, http.StatusBadRequest, answerJSON)

	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var user models.UserLogout
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

		if err := json.Unmarshal(body, &user); err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, http.StatusBadRequest, answerJSON)
			return
		}

		if err, code := h.services.User.Logout(user.UserID); err != nil {
			answerJSON = response.ErrorJSONCreator(err)
			response.WriteResponse(w, code, answerJSON)
			return
		}
		answerJSON = response.AnswerJSON{Data: "success", Code: http.StatusOK}
		response.WriteResponse(w, http.StatusOK, answerJSON)
		return

	default:
		errMessage := fmt.Sprintf("%s method not allowed", r.Method)
		errorJSON := &response.ErrorJson{Message: errMessage}
		answerJSON = response.AnswerJSON{Data: errorJSON, Code: http.StatusBadRequest}
		response.WriteResponse(w, http.StatusBadRequest, answerJSON)
		return
	}
}
