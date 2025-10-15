package utils

import "net/http"

type HTTPClient interface {
	Head(url string) (resp *http.Response, err error)
}
