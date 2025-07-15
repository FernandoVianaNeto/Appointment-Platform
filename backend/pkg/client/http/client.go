package http_client

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/adapters/$GOPACKAGE/mock_$GOFILE

import (
	configs "appointment-platform-backend-backend/cmd/config"
	"appointment-platform-backend-backend/internal/application/exceptions"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	baseUrl string
	client  *http.Client
	headers []OptionalHeaders
}

type OptionalHeaders struct {
	Key   string
	Value string
}

type HTTPClient interface {
	Get(ctx context.Context, endpointPath string, params url.Values) ([]byte, int, error)
	Post(ctx context.Context, endpoint string, params url.Values, body []byte) ([]byte, int, error)
	GetWithBody(ctx context.Context, endpointPath string, params url.Values, body []byte) ([]byte, int, error)
	Delete(ctx context.Context, endpoint string, params url.Values, body []byte) ([]byte, int, error)
	Do(ctx context.Context, method string, endpointPath string, params url.Values, body []byte) ([]byte, int, error)
	SetHeaders(headers ...OptionalHeaders)
}

var _ HTTPClient = &Client{}

func NewBaseClient(baseUrl string, client *http.Client, config *configs.ApplicationConfig, headers ...OptionalHeaders) HTTPClient {
	return &Client{
		baseUrl: baseUrl,
		client:  client,
		headers: headers,
	}
}

func (c *Client) SetHeaders(headers ...OptionalHeaders) {
	c.headers = headers
}

func (c *Client) Get(ctx context.Context, endpointPath string, params url.Values) ([]byte, int, error) {
	return c.Do(ctx, "GET", endpointPath, params, nil)
}

func (c *Client) Post(ctx context.Context, endpoint string, params url.Values, body []byte) ([]byte, int, error) {
	return c.Do(ctx, "POST", endpoint, params, body)
}

func (c *Client) GetWithBody(ctx context.Context, endpointPath string, params url.Values, body []byte) ([]byte, int, error) {
	return c.Do(ctx, "GET", endpointPath, params, body)
}

func (c *Client) Delete(ctx context.Context, endpoint string, params url.Values, body []byte) ([]byte, int, error) {
	return c.Do(ctx, "DELETE", endpoint, params, body)
}

func (c *Client) Do(ctx context.Context, method string, endpointPath string, params url.Values, body []byte) ([]byte, int, error) {
	u, _ := url.Parse(c.baseUrl)
	u.Path = path.Join(u.Path, endpointPath)

	if params != nil {
		u.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, http.StatusInternalServerError, exceptions.NewApplicationError(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "appointment-platform-backend-backend")

	for _, h := range c.headers {
		req.Header.Set(h.Key, h.Value)
	}

	resp, err := c.client.Do(req)
	s := req.URL.String()
	print(s)
	if err != nil {
		return nil, http.StatusInternalServerError, exceptions.NewApplicationError(err)
	}

	data, _ := io.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		return nil, http.StatusInternalServerError, exceptions.NewApplicationError(err)
	}

	return data, resp.StatusCode, nil
}

func HttpResponseError(code int, raw []byte) error {
	if code == http.StatusNotFound {
		return exceptions.NewNotFoundRestError("resource account not found")
	}

	return fmt.Errorf("HTTP status %d returns: %s", code, raw)
}

func JsonUnmarshalError(err error, raw []byte) error {
	return fmt.Errorf("unsmarshal error: %s. content: %s", err, raw)
}

func BindHttpResponse[V interface{}](target V, raw []byte, statusCode int) (V, error) {
	if statusCode != http.StatusOK {
		return target, HttpResponseError(statusCode, raw)
	}

	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	if err := decoder.Decode(&target); err != nil {
		return target, JsonUnmarshalError(err, raw)
	}

	return target, nil
}
