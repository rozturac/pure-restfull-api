package entity

type Config struct {
	key   string `json:"key"`
	value string `json:"value"`
}

func CreateConfig(key, value string) *Config {
	return &Config{
		key:   key,
		value: key,
	}
}