package controllers

import (
	"encoding/json"
	"github.com/rozturac/go-mediator"
	"net/http"
	"pure-restfull-api/api/configs"
	"pure-restfull-api/application/common"
	. "pure-restfull-api/application/create-config"
	. "pure-restfull-api/application/get-config"
)

type ConfigController struct {
	sender        mediator.Sender
	globalization *configs.Globalization
}

func NewConfigController(sender mediator.Sender, globalization *configs.Globalization) *ConfigController {
	return &ConfigController{
		sender:        sender,
		globalization: globalization,
	}
}

func (c *ConfigController) ServeHTTP() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			c.createConfig(writer, request)
		case http.MethodGet:
			c.getByKey(writer, request)
		}
	}
}

func (c *ConfigController) createConfig(writer http.ResponseWriter, request *http.Request) {
	var command *CreateConfigCommand

	if err := json.NewDecoder(request.Body).Decode(&command); err != nil {
		panic(common.UnExpectedCommand("RequestBody", command))
	}

	if _, err := c.sender.Send(newContext(c.globalization), command); err != nil {
		panic(err)
	}

	NoContent(writer, http.StatusCreated)
}

func (c *ConfigController) getByKey(writer http.ResponseWriter, request *http.Request) {
	var value string
	if value = request.FormValue("key"); len(value) == 0 {
		panic(common.NullOrEmptyArgumentError("key"))
	}

	result, err := c.sender.Send(newContext(c.globalization), &GetConfigQuery{
		Key: value,
	})

	if err != nil {
		panic(err)
	}

	JSON(writer, http.StatusOK, result)
}
