package main

import (
	"os"
	"pure-restfull-api/api"
	"pure-restfull-api/api/configs"
)

func main() {
	config, err := configs.LoadConfig("./api", os.Getenv("ENV"))
	if err != nil {
		panic(err)
	}

	api.Init(config)
}
