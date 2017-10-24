package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var Config struct {
	Debug  bool
	Listen string

	Postgres struct {
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		SSL      bool   `yaml:"ssl"`
	} `yaml:"postgres"`
}

func init() {

	f, err := os.Open("/home/khorevaa/Документы/github/go-web-v8storage/webui/config/config.yml")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	f.Close()

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}
}
