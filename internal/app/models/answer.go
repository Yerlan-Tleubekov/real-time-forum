package models

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

type Answer struct {
	Data  interface{} `json:"data"`
	Error Error       `json:"error"`
	Code  int         `json:"code"`
}

func (a *Answer) FillFields(data interface{}, code int, err error) {
	e := &Error{}
	if err != nil {
		e.Message = err.Error()
	}
	a.Data = data
	a.Code = code
	a.Error = *e

}

func (a *Answer) RespondJson(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(a)
	if err != nil {
		a.FillFields(nil, http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		a.RespondJson(w, r)
		return
	}

}
