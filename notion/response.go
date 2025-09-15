package notion

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type Response struct {
	Body   *bytes.Buffer
	Header http.Header
	Status int
}

const defaultMaxMemory = 32 << 20 // 32mb

func NewResponse(resp *http.Response) (r *Response, err error) {
	defer resp.Body.Close()

	limitReader := io.LimitReader(resp.Body, defaultMaxMemory)

	bodyBytes, err := io.ReadAll(limitReader)
	if err != nil {
		return nil, err
	}
	r = &Response{
		Body:   bytes.NewBuffer(bodyBytes),
		Header: resp.Header,
		Status: resp.StatusCode,
	}
	return r, nil
}

func (r *Response) Bind(v any) error {
	switch {
	case strings.HasPrefix(r.Header.Get("Content-Type"), "application/json"):
		fallthrough
	default:
		return json.Unmarshal(r.Body.Bytes(), v)
	}
}

func (r *Response) Ok() bool {
	return r.Status >= 200 && r.Status < 300
}

func (r *Response) LogValue() slog.Value {
	attrs := []slog.Attr{
		slog.Int("status", r.Status),
		slog.Any("header", r.headerLogValue()),
	}
	if r.Body != nil {
		attrs = append(attrs,
			slog.Any("body", r.bodyLogValue()),
		)
	}
	return slog.GroupValue(attrs...)
}

func (r *Response) bodyLogValue() (v slog.Value) {
	var err error
	defer func() {
		if err != nil {
			v = slog.StringValue(r.Body.String())
		}
	}()
	switch {
	case strings.HasPrefix(r.Header.Get("Content-Type"), "application/json"):
		var body map[string]any
		err = json.Unmarshal(r.Body.Bytes(), &body)
		if err != nil {
			return
		}
		v = slog.AnyValue(body)
	default:
		v = slog.StringValue(r.Body.String())
	}
	return
}

func (r *Response) headerLogValue() slog.Value {
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
