package request

import (
	"encoding/json"
	"net/url"
)

func WithMethod(method string) Option {
	return func(r *Request) error {
		r.Method = method
		return nil
	}
}

func WithUrl(u *url.URL) Option {
	return func(r *Request) error {
		r.Url = u
		return nil
	}
}

func WithRawUrl(rawUrl string) Option {
	return func(r *Request) error {
		u, err := url.Parse(rawUrl)
		if err != nil {
			return err
		}
		return WithUrl(u)(r)
	}
}

func WithHeader(key, value string) Option {
	return func(r *Request) error {
		r.Header.Add(key, value)
		return nil
	}
}

func WithBody(body []byte) Option {
	return func(r *Request) error {
		r.Body = body
		return nil
	}
}

func WithJsonBody(data any) Option {
	return func(r *Request) error {
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return WithBody(jsonBytes)(r)
	}
}
