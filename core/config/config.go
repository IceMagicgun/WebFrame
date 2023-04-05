package config

import (
	debugM "WebFrame/core/debug"
	"encoding/json"
	"io/ioutil"
	"os"
)

type ConfigStruct struct {
	Env    string
	OpenAI struct {
		OPENAI_API_KEY string
	}
	Proxy map[string]string
}

var (
	hasInitConfig = false
	Config        *ConfigStruct
)

func InitConfig(path string) {
	if hasInitConfig {
		return
	}
	initConfig(path)
}

func EnforceInitConfig(path string) {
	initConfig(path)
}

func initConfig(path string) {
	hasInitConfig = true
	Config = &ConfigStruct{}
	data, _ := ioutil.ReadFile("json/config.json")
	if path != "" {
		_, err := os.Stat(path)
		if err == nil {
			data, _ = ioutil.ReadFile(path)
		}
	}
	err := json.Unmarshal(data, &Config)
	if err != nil {
		debugM.Error(err.Error())
	}
}
