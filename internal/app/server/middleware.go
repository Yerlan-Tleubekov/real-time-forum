package server

import (
	"errors"
	"fmt"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/models"
	"net/http"
)

type Options struct {
	h http.Handler
}

func checkAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		answer := &models.Answer{}
		cookie, err := r.Cookie("forum")
		if err != nil || cookie == nil {
			answer.FillFields(nil, http.StatusUnauthorized, errors.New("unauthorized"))
			answer.RespondJson(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func setContentTypeMiddleware(next http.Handler, contType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contType)
		next.ServeHTTP(w, r)
	})
}

func checkMethodsMiddleware(next http.Handler, methods []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		answer := &models.Answer{}
		errText := "method %s not allowed"

		for _, v := range methods {
			if v == method {
				next.ServeHTTP(w, r)
				return
			}
		}

		answer.FillFields(nil, http.StatusBadRequest, errors.New(fmt.Sprintf(errText, method)))
		answer.RespondJson(w, r)
		return
	})
}

func Middlewares(handler http.Handler, methods []string, contType string, checkAuth bool) http.Handler {
	var h http.Handler
	if checkAuth {
		h = checkAuthMiddleware(handler)
	} else {
		h = handler
	}
	h = checkMethodsMiddleware(h, methods)
	h = setContentTypeMiddleware(h, contType)

	return h
}
