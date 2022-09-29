package settings

import (
	"encoding/json"
	"io/ioutil"
	"fmt"

	config "backengine/config"
)

var environments = map[string]string{
	"dev": "./auth/settings/dev.json",
}

type Settings struct {
	PrivateKeyPath string
	PublicKeyPath string
	JWTExpirationDelta int
}

var settings *Settings

func loadSettingsByEnv(env string) {
	fileName := environments[env]
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error while reading 'jwt' config file")
	}

	settings = &Settings{}
	jsonErr := json.Unmarshal(content, settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing 'JWT' config file", jsonErr)
	}
}

func Get() Settings {
	if settings == nil {
		env := config.Get().Env
		if env == "" {
			env = "dev"
		}
		loadSettingsByEnv(env)
	}
	return *settings
}