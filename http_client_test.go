package gUtils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestHttpClientDefaultHeader(t *testing.T) {
	var got http.Header
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Clone()
		_, _ = w.Write([]byte("ok"))
	}))
	defer srv.Close()

	// 兼容老项目：不传 option 时不注入任何额外请求头
	if _, err := NewHttpClient(time.Second).PostForm(srv.URL, "a=1"); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got.Get("X-Token"), "")

	h := NewHttpClient(time.Second, WithHeader("X-Token", "abc"), WithHeader("User-Agent", "gUtils-test"))
	if _, err := h.PosJson(srv.URL, `{"a":1}`); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got.Get("X-Token"), "abc")
	assert.Equal(t, got.Get("Content-Type"), "application/json")

	if _, err := h.Get(srv.URL, "a=1"); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got.Get("X-Token"), "abc")
}

func TestHttpClientHttp2Option(t *testing.T) {
	transport := func(h *HttpClient) *http.Transport {
		return h.client.Transport.(*http.Transport)
	}

	// 默认不设置 Transport，沿用标准库行为
	assert.Equal(t, NewHttpClientDefault().client.Transport == nil, true)

	h2 := NewHttpClientDefault(WithHttp2(true))
	assert.Equal(t, transport(h2).ForceAttemptHTTP2, true)
	assert.Equal(t, transport(h2).TLSNextProto == nil, true)

	h1 := NewHttpClientDefault(WithHttp2(false))
	assert.Equal(t, transport(h1).ForceAttemptHTTP2, false)
	assert.Equal(t, transport(h1).TLSNextProto != nil, true)
}
