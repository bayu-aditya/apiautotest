package rest

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/bayu-aditya/apiautotest/internal/core"
)

func (e *restEngine) Get(ctx context.Context, url string, props core.RestEngineGetInput) (*core.RestEngineGetOutput, error) {
	// define http transport
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// define http client
	client := &http.Client{
		Transport: transport,
	}

	// define http request
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
