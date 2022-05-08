package repository

import (
	"crypto/tls"
	"net/http"
	"time"
)

func Client() *http.Client {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := http.Client{
		Transport: &transport,
		Timeout:   10 * time.Second,
	}

	return &client
}
