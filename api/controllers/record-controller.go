package controllers

import (
	"github.com/rozturac/go-mediator"
	"pure-restfull-api/api/configs"
)

type RecordController struct {
	sender        mediator.Sender
	globalization *configs.Globalization
}

func NewRecordController(sender mediator.Sender, globalization *configs.Globalization) *ConfigController {
	return &ConfigController{
		sender:        sender,
		globalization: globalization,
	}
}