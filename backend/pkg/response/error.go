package response

import "net/http"

type ErrorJson struct {
	Message string `json:"message"`
}

func ErrorJSONCreator(err error) AnswerJSON {
	errorJSON := ErrorJson{Message: err.Error()}
	answerJSON := AnswerJSON{Data: errorJSON, Code: http.StatusBadRequest}
	return answerJSON
}

func ErrorJsonWriter(err error, code int, w http.ResponseWriter) {
	errorJSON := ErrorJSONCreator(err)
	answerJSON := AnswerJSON{Data: errorJSON, Code: http.StatusBadRequest}
	WriteResponse(w, code, answerJSON)
}
