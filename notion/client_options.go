package notion

import (
	"log/slog"
	"net/url"
)

type ClientOption func(c *Client) error

func ClientWithApiBaseUrl(u string) ClientOption {
	return func(c *Client) (err error) {
		c.apiBaseUrl, err = url.Parse(u)
		return err
	}
}

func ClientWithApiVersion(v string) ClientOption {
	return func(c *Client) error {
		c.apiVersion = v
		return nil
	}
}

func ClientWithLogger(logger *slog.Logger) ClientOption {
	return func(c *Client) error {
		c.logger = logger
		return nil
	}
}
