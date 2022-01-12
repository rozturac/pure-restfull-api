package controllers

import (
	"github.com/rozturac/go-mediator"
	"pure-restfull-api/api/configs"
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
