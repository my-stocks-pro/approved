package config

import (
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type TypeConfig struct {
	BaseURL string
	ApiURL  string
	Session string
	Token   string
	NewURL  string
}

func GetConfig() *TypeConfig {

	conf := &TypeConfig{}

	data, errReadFile := ioutil.ReadFile("config/approved-service.yaml")
	if errReadFile != nil {
		log.Fatalf("error: %v", errReadFile)
	}

	errYaml := yaml.Unmarshal(data, &conf)
	if errYaml != nil {
		log.Fatalf("error: %v", errYaml)
	}

	return conf
}
