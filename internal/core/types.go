package core

import "time"

type RestEngineGetInput struct{}

type RestEngineOutput struct {
	StatusCode int
	Duration   time.Duration
}
