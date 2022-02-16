package rest

import (
	"context"
	"net/http"

	"github.com/bayu-aditya/apiautotest/internal/core"
)

func (e *restEngine) Get(ctx context.Context, url string, props core.RestEngineGetInput) (*core.RestEngineGetOutput, error) {
	var client = new(http.Client)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	output := &core.RestEngineGetOutput{
		StatusCode: response.StatusCode,
	}
	return output, nil
}
