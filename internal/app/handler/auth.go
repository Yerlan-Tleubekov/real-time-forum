package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/pkg/response"
)

func (h *Handler) SignUp() http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var user models.User
			var answer models.Answer
			//var userPasswords models.UserPasswords

			body, err := ioutil.ReadAll(r.Body)

			if err != nil {
				answer.FillFields(nil, http.StatusBadRequest, err)
				answer.RespondJson(w, r)
				return
			}

			if err := json.Unmarshal(body, &user); err != nil {
				answer.FillFields(nil, http.StatusBadRequest, err)
				answer.RespondJson(w, r)

				return
			}

			if err := h.services.User.ComparePasswords(&user); err != nil {
				answer.FillFields(nil, http.StatusBadRequest, err)
				answer.RespondJson(w, r)
				return
			}

			if err, code := h.services.User.SignUp(&user); err != nil {
				answer.FillFields(nil, code, err)
				answer.RespondJson(w, r)

				return
			}

			successJSON := SignUpSuccess{UserID: user.ID}
			answer.FillFields(successJSON, http.StatusOK, nil)
			answer.RespondJson(w, r)

		})
}

func (h *Handler) SignIn() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var answer models.Answer
		var user models.UserSignIn

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			answer.FillFields(nil, http.StatusInternalServerError, err)
			answer.RespondJson(w, r)
			return
		}

		if err := json.Unmarshal(body, &user); err != nil {
			answer.FillFields(nil, http.StatusBadRequest, err)
			answer.RespondJson(w, r)
			return
		}

		session, err, code := h.services.User.SignIn(&user)
		if err != nil {
			answer.FillFields(nil, code, err)
			answer.RespondJson(w, r)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "forum",
			Path:    "/",
			Value:   session.Token,
			Expires: session.ExpTime,
		})

		successJSON := SignUpSuccess{UserID: session.UserID}
		answer.FillFields(successJSON, http.StatusOK, nil)
		answer.RespondJson(w, r)
	})
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
