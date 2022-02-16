package internal

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/bayu-aditya/apiautotest/internal/model"
)

func ReadInput(loc string) (*model.InputYaml, error) {
	yamlFile, err := ioutil.ReadFile(loc)
	if err != nil {
		return nil, err
	}

	var input model.InputYaml
	if err = yaml.Unmarshal(yamlFile, &input); err != nil {
		return nil, err
	}

	return &input, nil
}
