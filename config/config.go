package config

import (
	"encoding/json"
	"sync"
	"fmt"
	"io/ioutil"
)

type Config struct {
	AppName string `json: "appName"`
	URI struct {
		App	string `json:"api"`
	}	`json:"uri"`
	Port int `json:"port"`
	Env string `json:"env"`
	Version	string `json:"version"`
	Database	struct {
		Type string `json:"type"`
		URI       string `json:"uri"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Timeout   int64  `json:"timeout"`
		RunScript bool   `json:"runScript"`
	} `json:"database"`
	SessionTimeout int    `json:"sessionTimeout"` //specify session timeout in seconds
}

var (
	cfg Config
	once sync.Once
)

func Parse(file string) Config {
	once.Do(func() {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("error in reading config file", err);
		}
		if err := json.Unmarshal(data, &cfg); err != nil {
			fmt.Println("json parsing failed in config file", err)
		}
	})
	return cfg
}

func Get() Config {
	return cfg
}