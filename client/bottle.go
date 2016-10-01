package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// ShowBottlePath computes a request path to the show action of bottle.
func ShowBottlePath(bottleID int) string {
	return fmt.Sprintf("/bottles/%v", bottleID)
}

// 指定されたIDのボトルを取得する
func (c *Client) ShowBottle(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBottleRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBottleRequest create the request corresponding to the show action endpoint of the bottle resource.
func (c *Client) NewShowBottleRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
