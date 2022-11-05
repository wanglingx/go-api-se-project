package app

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ErrorCode struct {
	Status      string `yaml:"status"`
	MessageCode string `yaml:"message_code"`
	Th          string `yaml:"th"`
	En          string `yaml:"en"`
}

type ErrorMessage struct {
	BadRequest                ErrorCode `yaml:"bad_request"`
}

func InitErrorMessage(cv *Configs) *ErrorMessage {
	em := ErrorMessage{}

	file, err := os.Open(cv.ErrorPath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}

	err = yaml.Unmarshal(content, &em)
	if err != nil {
		panic(err.Error())
	}

	return &em
}