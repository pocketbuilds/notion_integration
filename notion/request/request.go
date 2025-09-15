package request

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	Method string
	Url    *url.URL
	Header http.Header
	Body   []byte
}

type Option func(r *Request) error

func New(opts ...Option) (*Request, error) {
	r := &Request{
		Header: make(http.Header),
	}
	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func (r *Request) LogValue() slog.Value {
	attrs := []slog.Attr{
		slog.String("method", r.Method),
		slog.String("url", r.Url.String()),
		slog.Any("header", r.headerLogValue()),
	}
	if r.Body != nil {
		attrs = append(attrs,
			slog.Any("body", r.bodyLogValue()),
		)
	}
	return slog.GroupValue(attrs...)
}

func (r *Request) bodyLogValue() (v slog.Value) {
	var err error
	defer func() {
		if err != nil {
			v = slog.StringValue(string(r.Body))
		}
	}()
	switch {
	case strings.HasPrefix(r.Header.Get("Content-Type"), "application/json"):
		var body map[string]any
		err = json.Unmarshal(r.Body, &body)
		if err != nil {
			return
		}
		v = slog.AnyValue(body)
	default:
		v = slog.StringValue(string(r.Body))
	}
	return
}

func (r *Request) headerLogValue() slog.Value {
	attrs := []slog.Attr{}
	for key, values := range r.Header {
		if len(values) == 1 {
			attrs = append(attrs, slog.String(key, values[0]))
		} else {
			attrs = append(attrs, slog.Any(key, values))
		}
	}
	return slog.GroupValue(attrs...)
}
