package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) HelloWorld(responseWriter http.ResponseWriter, request *http.Request) {
	WriteOkResponse(responseWriter, "Hello World!")
}



type ErrorResponseBody struct {
	Code string
	Message string
}

func WriteErrorResponse(writer http.ResponseWriter, status int, code string, message string) {
	writer.Header().Set("Content-Type", "application/json;charset=utf-8")
	writer.WriteHeader(status)
	errorResponseBody := ErrorResponseBody{
		Code:    code,
		Message: message,
	}
	if encodingError := json.NewEncoder(writer).Encode(errorResponseBody); encodingError != nil {
		fmt.Printf("Aie")
	}
}

func WriteOkResponse(writer http.ResponseWriter, body interface{}) {
	writeSuccessWithContent(writer, body, http.StatusOK)
}

func writeSuccessWithContent(writer http.ResponseWriter, body interface{}, status int) {
	writer.Header().Add("Content-Type", "application/json;charset=utf-8")
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(body)
	if err != nil {
		WriteErrorResponse(writer, http.StatusInternalServerError, "InternalError", "fail to encode body object")
	}
}