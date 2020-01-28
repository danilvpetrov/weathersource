package collection

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const (
	// BaseURL is the base URL for the HTTP client to use.
	BaseURL = "https://api.darksky.net/forecast"
)

// ClientOptions is the set of options that can be passed to the client.
type ClientOptions struct {
	Client    *http.Client
	Longitude float64
	Latitude  float64
	Units     string
	Exclude   string
	Extend    string
}

// DefaultClientOptions returns the ClientOptions filled with default values.
func DefaultClientOptions() *ClientOptions {
	return &ClientOptions{
		Client:    http.DefaultClient,
		Longitude: 0,
		Latitude:  0,
		Units:     "ca",
		Exclude:   "",
		Extend:    "",
	}
}

// ClientOption is an option that can be passed to NewClient function.
type ClientOption func(*ClientOptions)

// WithLatAndLong is an option to add latitude and longitude to the HTTP client.
func WithLatAndLong(lat float64, lon float64) ClientOption {
	return func(options *ClientOptions) {
		options.Latitude = lat
		options.Longitude = lon
	}
}

// TO-DO: add more function options here

// Client is the HTTP client used to request weather information from
type Client struct {
	options *ClientOptions
	token   string
}

// NewClient returns a new instance of Client.
func NewClient(token string, opts ...ClientOption) *Client {
	options := DefaultClientOptions()

	for _, opt := range opts {
		opt(options)
	}

	return &Client{
		options: options,
		token:   token,
	}
}

// Forecast receives forecast from https://api.darksky.net.
func (c *Client) Forecast(ctx context.Context) (*Response, error) {
	u := c.makeURLFromOpts()
	fr := &Response{}
	err := c.Get(ctx, u, fr)

	return fr, err
}

// Get performs a GET method request at the given url and JSON-decodes the
// received payload into data. data has to be assignable in order for decoding
// to work.
func (c *Client) Get(ctx context.Context, u *url.URL, data interface{}) error {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := c.options.Client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(data)
}

func (c *Client) makeURLFromOpts() *url.URL {
	u := mustParse(url.Parse(BaseURL))

	u.Path = path.Join(u.Path, c.token)
	u.Path = path.Join(
		u.Path,
		fmt.Sprintf(
			"%g,%g",
			c.options.Latitude,
			c.options.Longitude,
		),
	)

	q := u.Query()

	if c.options.Units != "" {
		q.Add("units", c.options.Units)
	}

	if c.options.Exclude != "" {
		q.Add("exclude", c.options.Exclude)
	}

	if c.options.Extend != "" {
		q.Add("extend", c.options.Extend)
	}

	u.RawQuery = q.Encode()

	return u
}

func mustParse(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	return u
}
