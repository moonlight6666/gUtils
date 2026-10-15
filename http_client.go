package gUtils

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpClient struct {
	client  *http.Client
	headers http.Header
}

// HttpClientOption 在新建客户端时追加配置，不传任何 option 时行为与旧版完全一致。
type HttpClientOption func(*HttpClient)

func NewHttpClientDefault(opts ...HttpClientOption) *HttpClient {
	return NewHttpClient(time.Second*120, opts...)
}

func NewHttpClient(timeout time.Duration, opts ...HttpClientOption) *HttpClient {
	h := &HttpClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
	for _, opt := range opts {
		if opt != nil {
			opt(h)
		}
	}
	return h
}

// WithHttp2 设置是否使用 http2；不调用则跟随标准库默认（https 自动协商 h2）。
func WithHttp2(enable bool) HttpClientOption {
	return func(h *HttpClient) {
		tr, ok := http.DefaultTransport.(*http.Transport)
		var t *http.Transport
		if ok {
			t = tr.Clone()
		} else {
			t = &http.Transport{}
		}
		if enable {
			t.ForceAttemptHTTP2 = true
			t.TLSNextProto = nil
		} else {
			// 非 nil 空 map 是标准库中禁用 http2 的方式
			t.ForceAttemptHTTP2 = false
			t.TLSNextProto = make(map[string]func(authority string, c *tls.Conn) http.RoundTripper)
		}
		h.client.Transport = t
	}
}

// WithHeader 设置默认请求头，每次请求自动注入；同名的单次请求头优先。
func WithHeader(key, value string) HttpClientOption {
	return func(h *HttpClient) {
		if h.headers == nil {
			h.headers = make(http.Header)
		}
		h.headers.Add(key, value)
	}
}

func (h *HttpClient) do(method, u, contentType string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	for key, values := range h.headers {
		if req.Header.Get(key) == "" {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	return io.ReadAll(resp.Body)
}

func (h *HttpClient) GetWithStruct(u string, data any) (rs []byte, err error) {
	form, err := StructToForm(data, "form")
	if err != nil {
		return nil, err
	}
	return h.Get(u, form)
}

func (h *HttpClient) Get(u string, form string) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(u)
	if err != nil {
		return nil, err
	}
	Url.RawQuery = form

	return h.do(http.MethodGet, Url.String(), "", nil)
}

func (h *HttpClient) Post(url string, contentType string, data string) ([]byte, error) {
	return h.do(http.MethodPost, url, contentType, strings.NewReader(data))
}

func (h *HttpClient) PostForm(url string, form string) ([]byte, error) {
	return h.Post(url, "application/x-www-form-urlencoded", form)
}

func (h *HttpClient) PosJson(url string, data string) ([]byte, error) {
	return h.Post(url, "application/json", data)
}

func (h *HttpClient) PosFormWithStruct(url string, data any) ([]byte, error) {
	form, err := StructToForm(data, "form")
	if err != nil {
		return nil, err
	}
	return h.PostForm(url, form)
}

func (h *HttpClient) PosJsonWithStruct(url string, data any) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return h.PosJson(url, string(b))
}
