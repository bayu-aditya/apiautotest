package input

import "github.com/bayu-aditya/apiautotest/internal/core"

func New() core.IInputConfiguration {
	return &inputImpl{}
}
