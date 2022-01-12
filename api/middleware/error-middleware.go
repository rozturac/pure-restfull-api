package middleware

import (
	"github.com/rozturac/cerror"
	"net/http"
	"pure-restfull-api/api/controllers"
)

func ErrorMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			response := mapper(http.StatusText(http.StatusInternalServerError))
			httpStatusCode := http.StatusInternalServerError
			if recover := recover(); recover != nil {
				if err, ok := recover.(error); ok {
					response = mapper(err.Error())
					if customError, ok := err.(cerror.Error); ok {
						httpStatusCode = customError.HttpStatusCode()
					}
				}

				controllers.JSON(writer, httpStatusCode, response)
			}
		}()

		next(writer, request)
	}
}

func mapper(message string) interface{} {
	type ConsistentErrorMessage struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	}

	return &ConsistentErrorMessage{
		Code:    -1,
		Message: message,
	}
}
