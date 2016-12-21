package govrealize

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

var (
	mux *http.ServeMux

	client *Client

	server *httptest.Server
)

// TokenSource represents
type TokenSource struct {
	AccessToken string
	Expiry      time.Time
}

// Token represents
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
		Expiry:      t.Expiry,
	}
	return token, nil
}

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	tokenSource := &TokenSource{
		AccessToken: "token.ID",
		Expiry:      time.Now(),
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)

	client = NewClient(oauthClient)

	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

type values map[string]string

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testClientServices(t *testing.T, c *Client) {
	services := []string{
		"Machine",
	}

	cp := reflect.ValueOf(c)
	cv := reflect.Indirect(cp)

	for _, s := range services {
		if cv.FieldByName(s).IsNil() {
			t.Errorf("c.%s shouldn't be nil", s)
		}
	}
}

func testClientDefaults(t *testing.T, c *Client) {
	testClientServices(t, c)
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)
	testClientDefaults(t, c)
}

func TestNew(t *testing.T) {
	c, err := New(nil)

	if err != nil {
		t.Fatalf("New(): %v", err)
	}
	testClientDefaults(t, c)
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)
	u, err := url.Parse("https://qwertyuiop.url/")
	if err != nil {
		t.Fatalf("NewRequest(): %v", err)
	}
	c.BaseURL = u

	inURL, outURL := "/foo", c.BaseURL.String()+"foo"
	inBody, outBody := &MachineCreateRequest{
		Type: "CatalogItemRequest",
		CatalogItemRef: MachineCreateRequestCatalogItemRef{
			ID: "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
		},
		Organization: MachineCreateRequestOrganization{
			TenantRef:    "vsphere.local",
			SubtenantRef: "f04f060d-73be-48a3-b82c-20cb98efd2d2",
		},
		State:         "SUBMITTED",
		RequestNumber: 0,
	},
		`{"@type":"CatalogItemRequest","catalogItemRef":{"id":"c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6"},"organization":{"tenantRef":"vsphere.local","subtenantRef":"f04f060d-73be-48a3-b82c-20cb98efd2d2"},"state":"SUBMITTED","requestNumber":0,"requestData":{"entries":null}}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v)Body = %v, expected %v", inBody, string(body), outBody)
	}
}

func TestNewRequest_withUserData(t *testing.T) {
	c := NewClient(nil)
	u, err := url.Parse("https://qwertyuiop.url/")
	if err != nil {
		t.Fatalf("NewRequest(): %v", err)
	}
	c.BaseURL = u

	inURL, outURL := "/foo", c.BaseURL.String()+"foo"
	inBody, outBody := &MachineCreateRequest{
		Type: "CatalogItemRequest",
		CatalogItemRef: MachineCreateRequestCatalogItemRef{
			ID: "c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6",
		},
		Organization: MachineCreateRequestOrganization{
			TenantRef:    "vsphere.local",
			SubtenantRef: "f04f060d-73be-48a3-b82c-20cb98efd2d2",
		},
		State:         "SUBMITTED",
		RequestNumber: 0,
	},
		`{"@type":"CatalogItemRequest","catalogItemRef":{"id":"c94fa0c3-4aed-43ce-b7a6-4163a07e4cd6"},"organization":{"tenantRef":"vsphere.local","subtenantRef":"f04f060d-73be-48a3-b82c-20cb98efd2d2"},"state":"SUBMITTED","requestNumber":0,"requestData":{"entries":null}}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v)Body = %v, expected %v", inBody, string(body), outBody)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewRequest("GET", ":", nil)
	testURLParseError(t, err)
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(foo)
	_, err := client.Do(req, body)
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	expected := &foo{"a"}
	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Response body = %v, expected %v", body, expected)
	}
}

func TestDo_httpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Error("Expected HTTP 400 error.")
	}
}

// Test handling of an error caused by the internal http client's Do()
// function.
func TestDo_redirectLoop(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected a URL error; got %#v.", err)
	}
}

func TestCheckResponse(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(strings.NewReader(`{"errors": [{
    		"code": 10101,
    		"message": "Invalid argument.",
    		"systemMessage": "Request body or some of its properties expected for post or put request but was empty.",
    		"moreInfoUrl": null
    	}]}`)),
	}
	err := CheckResponse(res).(*ErrorResponse)

	if err == nil {
		t.Fatalf("Expected error response.")
	}

	error := ErrorResponseError{
		Code:          10101,
		Message:       "Invalid argument.",
		SystemMessage: "Request body or some of its properties expected for post or put request but was empty.",
	}

	errors := []ErrorResponseError{}

	errors = append(errors, error)

	expected := &ErrorResponse{
		Response: res,
		Errors:   errors,
	}
	if !reflect.DeepEqual(err, expected) {
		t.Errorf("Error = %#v, expected %#v", err, expected)
	}
}

// ensure that we properly handle API errors that do not contain a response
// body
func TestCheckResponse_noBody(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}
	err := CheckResponse(res).(*ErrorResponse)

	if err == nil {
		t.Errorf("Expected error response.")
	}

	expected := &ErrorResponse{
		Response: res,
	}
	if !reflect.DeepEqual(err, expected) {
		t.Errorf("Error = %#v, expected %#v", err, expected)
	}
}

func TestErrorResponse_Error(t *testing.T) {

	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(strings.NewReader(`{"errors": [{
            "code": 10101,
            "message": "Invalid argument.",
            "systemMessage": "Request body or some of its properties expected for post or put request but was empty.",
            "moreInfoUrl": null
        }]}`)),
	}
	err := CheckResponse(res).(*ErrorResponse)

	if err == nil {
		t.Fatalf("Expected error response.")
	}

	error := ErrorResponseError{
		Code:          10101,
		Message:       "Invalid argument.",
		SystemMessage: "Request body or some of its properties expected for post or put request but was empty.",
	}

	errors := []ErrorResponseError{}

	errors = append(errors, error)

	expected := &ErrorResponse{
		Response: res,
		Errors:   errors,
	}

	if expected.Error() == "" {
		t.Errorf("Expected non-empty ErrorResponse.Error()")
	}
}
