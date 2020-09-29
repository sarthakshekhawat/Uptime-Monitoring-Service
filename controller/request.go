package controller

import (
	"net/http"
	"time"

	"github.com/gojektech/heimdall/httpclient"
)

// RequestInterface is the interface for all the fuctions which makes the http request
type RequestInterface interface {
	httpRequest(data DataBase) (*http.Response, error)
}

// RequestReceiver is the receiver type for the functions which makes the http request
type RequestReceiver struct{}

var requestRepo RequestInterface

// AssignRequestRepo assigns the value to the requestRepo
func AssignRequestRepo(ri RequestInterface) {
	requestRepo = ri
}

func (rr *RequestReceiver) httpRequest(data DataBase) (*http.Response, error) {
	timeout := time.Duration(data.CrawlTimeout) * time.Second
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	return client.Get(data.URL, nil)
}
