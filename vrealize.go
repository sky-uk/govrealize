package govrealize

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	version   = "6.2"
	mediaType = "application/json"
)

// Client manages communication with the vRealize API 6.2
type Client struct {
	// HTTP client used to communicate with the vRealize API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// Username for client
	Username string

	// Password for client
	Password string

	// Tenant for client
	Tenant string

	// Services used for communicating with the API
	Machine MachineService

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	Limit int `url:"limit,omitempty"`
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// Response is a vRealize response. This wraps the standard http.Response returned from vRealize.
type Response struct {
	*http.Response

	Links []Link
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	Errors []ErrorResponseError `json:"errors"`
}

// ErrorResponseError represents an API error
type ErrorResponseError struct {
	Code          int         `json:"code"`
	Message       string      `json:"message"`
	SystemMessage string      `json:"systemMessage"`
	MoreInfoURL   interface{} `json:"moreInfoUrl"`
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	origURL, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	origValues := origURL.Query()

	newValues, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	for k, v := range newValues {
		origValues[k] = v
	}

	origURL.RawQuery = origValues.Encode()
	return origURL.String(), nil
}

// NewClient returns a new vRealize API client.
func NewClient(httpClient *http.Client) *Client {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{client: httpClient}

	c.Machine = &MachineServiceOp{client: c}

	return c
}

// ClientOpt are options for New.
type ClientOpt func(*Client) error

// New returns a new vRealize API client instance.
func New(httpClient *http.Client, opts ...ClientOpt) (*Client, error) {
	c := NewClient(httpClient)
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// SetBaseURL is a client option for setting the base URL.
func SetBaseURL(bu string) ClientOpt {
	return func(c *Client) error {
		u, err := url.Parse(bu)
		if err != nil {
			return err
		}

		c.BaseURL = u
		return nil
	}
}

// SetUsername is a client option for setting the Username.
func SetUsername(username string) ClientOpt {
	return func(c *Client) error {
		c.Username = fmt.Sprintf("%s+%s", username, c.Username)
		return nil
	}
}

// SetPassword is a client option for setting the Password.
func SetPassword(password string) ClientOpt {
	return func(c *Client) error {
		c.Password = fmt.Sprintf("%s+%s", password, c.Password)
		return nil
	}
}

// SetTenant is a client option for setting the Tenant.
func SetTenant(tenant string) ClientOpt {
	return func(c *Client) error {
		c.Tenant = fmt.Sprintf("%s+%s", tenant, c.Tenant)
		return nil
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		er := json.NewEncoder(buf).Encode(body)
		if er != nil {
			return nil, er
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)

	return req, nil
}

// OnRequestCompleted sets the DO API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

// newResponse creates a new Response for the provided http.Response
func newResponse(r *http.Response) *Response {
	response := Response{Response: r}

	return &response
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, resp)
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return response, err
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode)
}

// CheckResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}

	return errorResponse
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	p := new(string)
	*p = v
	return p
}

// Int is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it, but unlike Int32
// its argument value is an int.
func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

// StreamToString converts a reader to a string
func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(stream)
	return buf.String()
}
