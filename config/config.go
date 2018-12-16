package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Debug       bool          `json:"debug"`
	TokenConfig []TokenConfig `json:"token"`
}

type TokenConfig struct {
	Name     string       `json:"name"`
	Url      string       `json:"url"`
	Assemble string       `json:"assemble"`
	Parser   ParserConfig `json:"parser"`
}

type ParserConfig struct {
	Method string `json:"method"`
	Route  string `json:"route"`
}

func NewConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
