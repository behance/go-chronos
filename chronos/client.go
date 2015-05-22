package chronos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Constants to represent HTTP verbs
const (
	HTTPGet    = "GET"
	HTTPPut    = "PUT"
	HTTPDelete = "DELETE"
	HTTPPost   = "POST"
)

// Chronos is a client that can interact with the chronos API
type Chronos interface {
	Jobs() (*Jobs, error)
	DeleteJob(name string) error
}

// A Client can make http requests
type Client struct {
	config Config
	http   *http.Client
}

// NewClient returns a new chronos client, initialzed with the provided config
func NewClient(config Config) (Chronos, error) {
	client := new(Client)

	client.config = config

	client.http = &http.Client{
		Timeout: (time.Duration(config.RequestTimeout) * time.Second),
	}
	return client, nil
}

func (client *Client) apiGet(uri string, result interface{}) error {
	_, err := client.apiCall(HTTPGet, uri, "", result)
	return err
}

func (client *Client) apiDelete(uri string, result interface{}) error {
	_, err := client.apiCall(HTTPDelete, uri, "", result)
	return err
}

func (client *Client) apiCall(method, uri, body string, result interface{}) (int, error) {
	status, response, err := client.httpCall(method, uri, body)

	if err != nil {
		return 0, err
	}

	err = json.NewDecoder(response.Body).Decode(result)

	if err != nil {
		return status, err
	}

	// TODO: Handle error status codes
	return status, nil
}

// TODO: think about pulling out a Request struct/object/thing
func (client *Client) applyRequestHeaders(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
}

func (client *Client) newRequest(method, uri, body string) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", client.config.URL, uri)
	request, err := http.NewRequest(method, url, strings.NewReader(body))

	if err != nil {
		return nil, err
	}

	client.applyRequestHeaders(request)
	return request, nil
}

func (client *Client) httpCall(method, uri, body string) (int, *http.Response, error) {
	request, err := client.newRequest(method, uri, body)

	if err != nil {
		return 0, nil, err
	}

	response, err := client.http.Do(request)

	if err != nil {
		return 0, nil, err
	}

	return response.StatusCode, response, nil
}

// TODO: this better
func (client *Client) log(message string, args ...interface{}) {
	fmt.Printf(message+"\n", args...)
}
