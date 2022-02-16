package core

import "github.com/bayu-aditya/apiautotest/internal/model"

type IInputConfiguration interface {
	ReadYamlFromFs(fsLoc string) (*model.InputYaml, error)
}
