package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"pure-restfull-api/api/configs"
	"pure-restfull-api/application/common"
)

type BaseController interface {
	ServeHTTP() http.HandlerFunc
}

func newContext(globalization *configs.Globalization) context.Context {
	ctx := context.WithValue(context.Background(), common.CorrelationId, uuid.NewString())
	return context.WithValue(ctx, common.Location, globalization.Location)
}

func NoContent(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
