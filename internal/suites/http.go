package suites

import (
	"crypto/tls"
	"net/http"
)

// NewHTTPClient create a new client skipping TLS verification and not redirecting.
func NewHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}
