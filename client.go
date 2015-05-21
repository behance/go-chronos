package chronos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	HTTP_GET    = "GET"
	HTTP_PUT    = "PUT"
	HTTP_DELETE = "DELETE"
	HTTP_POST   = "POST"
)

type ChronosClient interface {
	Jobs() (*Jobs, error)
}

type Client struct {
	config Config
	http   *http.Client
}

func NewClient(config Config) (ChronosClient, error) {
	client := new(Client)

	client.config = config

	client.http = &http.Client{
		Timeout: (time.Duration(config.RequestTimeout) * time.Second),
	}
	return client, nil
}

func (client *Client) apiGet(uri string, result interface{}) error {
	_, error := client.apiCall(HTTP_GET, uri, "", result)
	return error
}

func (client *Client) apiCall(method, uri, body string, result interface{}) (int, error) {
	client.log("apiCall() method: %s, uri: %s, body: %s", method, uri, body)

	if status, response, err := client.httpCall(method, uri, body); err != nil {
		return 0, err
	} else {
		if err := json.NewDecoder(response.Body).Decode(result); err != nil {
			return status, err
		}
		// TODO: Handle error status codes
		return status, nil
	}
}

// TODO: think about pulling out a Request struct/object/thing
func (client *Client) applyRequestHeaders(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
}

func (client *Client) newRequest(method, uri, body string) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", client.config.URL, uri)
	if request, err := http.NewRequest(method, url, strings.NewReader(body)); err != nil {
		return nil, err
	} else {
		client.applyRequestHeaders(request)
		return request, nil
	}
}

func (client *Client) httpCall(method, uri, body string) (int, *http.Response, error) {
	if request, err := client.newRequest(method, uri, body); err != nil {
		return 0, nil, err
	} else {
		if response, err := client.http.Do(request); err != nil {
			return 0, nil, err
		} else {
			return response.StatusCode, response, nil
		}
	}
}

// TODO: this better
func (client *Client) log(message string, args ...interface{}) {
	fmt.Printf(message+"\n", args...)
}
