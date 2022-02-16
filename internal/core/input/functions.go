package input

import (
	"io/ioutil"

	"github.com/bayu-aditya/apiautotest/internal/model"
	"gopkg.in/yaml.v2"
)

func (i *inputImpl) ReadYamlFromFs(fsLoc string) (*model.InputYaml, error) {
	yamlFile, err := ioutil.ReadFile(fsLoc)
	if err != nil {
		return nil, err
	}

	var input model.InputYaml
	if err = yaml.Unmarshal(yamlFile, &input); err != nil {
		return nil, err
	}

	return &input, nil
}
