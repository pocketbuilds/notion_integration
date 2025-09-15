package notion

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/pocketbuilds/notion_integration/notion/request"
)

type Client struct {
	apiBaseUrl *url.URL
	apiVersion string
	apiKey     string
	logger     *slog.Logger
}

func NewClient(key string, opts ...ClientOption) (client *Client, err error) {
	client = &Client{
		apiKey:     key,
		apiVersion: "2022-06-28",
	}
	client.apiBaseUrl, err = url.Parse("https://api.notion.com/v1")
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		err = opt(client)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

func (c *Client) NewRequest(opts ...request.Option) (*request.Request, error) {
	return request.New(
		append([]request.Option{
			request.WithHeader("Authorization", "Bearer "+c.apiKey),
			request.WithHeader("Notion-Version", c.apiVersion),
		}, opts...)...,
	)
}

func (c *Client) DoRequest(req *request.Request) (*Response, error) {
	stdReq, err := http.NewRequest(req.Method, req.Url.String(), bytes.NewReader(req.Body))
	if err != nil {
		return nil, err
	}
	for key, values := range req.Header {
		for _, v := range values {
			stdReq.Header.Add(key, v)
		}
	}
	stdRes, err := http.DefaultClient.Do(stdReq)
	if err != nil {
		return nil, err
	}
	res, err := NewResponse(stdRes)
	if err != nil {
		return nil, err
	}
	if !res.Ok() {
		apiError := &Error{}
		if err := res.Bind(apiError); err != nil {
			return nil, err
		}
		c.log(req, res, apiError)
		return nil, apiError
	} else {
		c.log(req, res, nil)
		return res, nil
	}
}

func (c *Client) log(req *request.Request, res *Response, err *Error) {
	if c.logger == nil {
		return
	}
	msg := fmt.Sprintf("%s %s", req.Method, req.Url.String())
	attrs := []any{
		"request", req,
		"response", res,
	}
	if err != nil {
		attrs = append(attrs,
			slog.String("error", err.Message),
			// TODO: better details
			slog.Any("details", err),
		)
		c.logger.Error(msg, attrs...)
	} else {
		c.logger.Info(msg, attrs...)
	}
}

func (c *Client) Get(path string, morePath ...string) (*Response, error) {
	u := c.apiBaseUrl.JoinPath(path)
	for _, frag := range morePath {
		u = u.JoinPath(frag)
	}

	req, err := c.NewRequest(
		request.WithMethod(http.MethodGet),
		request.WithUrl(u),
	)
	if err != nil {
		return nil, err
	}

	return c.DoRequest(req)
}

func (c *Client) Post(body any, path string, morePath ...string) (*Response, error) {
	u := c.apiBaseUrl.JoinPath(path)
	for _, frag := range morePath {
		u = u.JoinPath(frag)
	}

	req, err := c.NewRequest(
		request.WithMethod(http.MethodPost),
		request.WithUrl(u),
		request.WithJsonBody(body),
		request.WithHeader("Content-Type", "application/json"),
	)
	if err != nil {
		return nil, err
	}

	return c.DoRequest(req)
}

func (c *Client) Patch(body any, path string, morePath ...string) (*Response, error) {
	u := c.apiBaseUrl.JoinPath(path)
	for _, frag := range morePath {
		u = u.JoinPath(frag)
	}

	req, err := c.NewRequest(
		request.WithMethod(http.MethodPatch),
		request.WithUrl(u),
		request.WithJsonBody(body),
		request.WithHeader("Content-Type", "application/json"),
	)
	if err != nil {
		return nil, err
	}

	return c.DoRequest(req)
}
