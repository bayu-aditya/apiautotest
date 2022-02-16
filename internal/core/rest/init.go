package rest

import "github.com/bayu-aditya/apiautotest/internal/core"

func New() core.IRestEngineCore {
	return &restEngine{}
}
