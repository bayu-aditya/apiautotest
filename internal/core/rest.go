package core

import "context"

type IRestEngineCore interface {
	Get(ctx context.Context, url string, props RestEngineGetInput) (*RestEngineOutput, error)
}
