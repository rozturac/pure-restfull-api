package entity

import "encoding/json"

type Config struct {
	key   string `json:"key"`
	value string `json:"value"`
}

func CreateConfig(key, value string) *Config {
	return &Config{
		key:   key,
		value: value,
	}
}

func (c *Config) GetKey() string {
	return c.key
}

func (c *Config) GetValue() string {
	return c.value
}

func (c *Config) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{
		Key:   c.key,
		Value: c.value,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}
