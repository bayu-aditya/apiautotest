package rest

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/bayu-aditya/apiautotest/internal/core"
)

func (e *restEngine) Get(ctx context.Context, url string, props core.RestEngineGetInput) (*core.RestEngineOutput, error) {
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

	startTime := time.Now()
	response, err := client.Do(request)
	durationTime := time.Since(startTime)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	output := &core.RestEngineOutput{
		StatusCode: response.StatusCode,
		Duration:   durationTime,
	}
	return output, nil
}
