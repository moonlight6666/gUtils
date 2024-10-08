package gUtils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClientDefault() *HttpClient {
	return NewHttpClient(time.Second * 120)
}

func NewHttpClient(timeout time.Duration) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (h *HttpClient) Get(u string, data string) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(u)
	if err != nil {
		return nil, err
	}
	Url.RawQuery = data

	resp, err := h.client.Get(Url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	return io.ReadAll(resp.Body)
}

func (h *HttpClient) Post(url string, contentType string, data string) ([]byte, error) {
	resp, err := h.client.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	return io.ReadAll(resp.Body)
}

func (h *HttpClient) PostForm(url string, data string) ([]byte, error) {
	return h.Post(url, "application/x-www-form-urlencoded", data)
}

func (h *HttpClient) PosJson(url string, data string) ([]byte, error) {
	return h.Post(url, "application/json", data)
}

func (h *HttpClient) PosJsonWithStruct(url string, data any) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return h.PosJson(url, string(b))
}
