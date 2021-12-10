package http

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

type Provider struct {
	http *http.Client
}

func New() *Provider {
	return &Provider{http: http.DefaultClient}
}

func (p *Provider) Get(url string, answer interface{}) error {
	if err := p.do(http.MethodGet, url, nil, answer); err != nil {
		return err
	}

	return nil
}

func (p *Provider) do(method, url string, body []byte, answer interface{}) error {
	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("error request from provider: %w", err)
	}

	switch method {
	case http.MethodPost, http.MethodPut:
		request.Header.Set("Content-Type", "application/json")
	}

	req, err := p.http.Do(request)
	if err != nil {
		return fmt.Errorf("error request from client provider: %w", err)
	}

	if req.StatusCode >= 300 {
		return fmt.Errorf("error request, statuscode: %d", req.StatusCode)
	}

	defer func() {
		if err = req.Body.Close(); err != nil {
			return
		}
	}()

	body, err = io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("error read body from provider: %w", err)
	}

	switch answer.(type) {
	case *string:
		*answer.(*string) = string(body)
	default:
		if err = jsoniter.Unmarshal(body, answer); err != nil {
			return fmt.Errorf("error unmashal json from provider: %w", err)
		}
	}

	return nil
}
