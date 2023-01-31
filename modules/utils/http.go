package utils

import (
	"fmt"
	"net/http"
	"time"
)

func HttpRequest(method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Printf("could not create request %s\n", err)
	}

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("error making http request %s\n", err)
	}

	return res, err
}
