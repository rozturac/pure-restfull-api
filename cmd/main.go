package main

import (
	"fmt"
	"github.com/rozturac/cerror"
	"os"
	"pure-restfull-api/api"
	"pure-restfull-api/api/configs"
)

func main() {

	defer func() {
		if recover := recover(); recover != nil {
			if err, ok := recover.(cerror.Error); ok {
				fmt.Println(fmt.Sprintf("[%s] %s", err.ErrorCode(), err.ErrorWithTrace()))
			} else {
				fmt.Println(fmt.Sprintf("%v", recover))
			}
		}
	}()

	config, err := configs.LoadConfig("./api", os.Getenv("ENV"))
	if err != nil {
		panic(err)
	}

	api.Init(config)
}
