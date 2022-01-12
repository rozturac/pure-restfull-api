package controllers

import (
	"encoding/json"
	"github.com/rozturac/go-mediator"
	"net/http"
	"pure-restfull-api/api/configs"
	"pure-restfull-api/application/common"
	. "pure-restfull-api/application/get-records"
)

type RecordController struct {
	sender        mediator.Sender
	globalization *configs.Globalization
}

func NewRecordController(sender mediator.Sender, globalization *configs.Globalization) *RecordController {
	return &RecordController{
		sender:        sender,
		globalization: globalization,
	}
}

func (r RecordController) ServeHTTP() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			r.getRecordsByFilter(writer, request)
		default:
			NoContent(writer, http.StatusNotFound)
		}
	}
}

func (r *RecordController) getRecordsByFilter(writer http.ResponseWriter, request *http.Request) {
	var command *GetRecordsByTimeAndCountRangeQuery

	if err := json.NewDecoder(request.Body).Decode(&command); err != nil {
		panic(common.UnExpectedCommand("RequestBody", command))
	}

	result, err := r.sender.Send(newContext(r.globalization), command)
	if err != nil {
		panic(err)
	}

	JSON(writer, http.StatusOK, result)
}
