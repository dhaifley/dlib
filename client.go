package dlib

import "net/http"

// ClientExecutor types are able to execute HTTP requests.
type ClientExecutor interface {
	Do(req *http.Request) (*http.Response, error)
}
