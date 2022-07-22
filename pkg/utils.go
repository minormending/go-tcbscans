package pkg

import (
	"io"
	"net/http"
	"os"
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

// savePage saves the response body to the given path.
func savePage(url string, filename string) error {
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()

		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}

		return nil
	}
	return err
}
