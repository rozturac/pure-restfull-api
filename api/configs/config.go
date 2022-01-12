package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Host          *Host
	MongoDB       *MongoDB
	Globalization *Globalization
}

type MongoDB struct {
	URI      string
	Database string
}

type Host struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

type Globalization struct {
	Location string
}

func LoadConfig(path, env string) (config *Config, err error) {

	filePath := fmt.Sprintf("%s/appsettings.%s.json", path, env)
	viper.SetConfigFile(filePath)

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
