package pkg

import (
	"io"
	"net/http"
)

// getPage returns the html page for the given url.
func getPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err == nil {
			return string(body), nil
		}
	}
	return "", err
}
