package pkg

import (
	"errors"
	"io"
	"net/http"
	"regexp"
)

type Series struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func GetSeries() ([]Series, error) {
	page, err := getSeriesPage()
	if err != nil {
		return nil, err
	}

	var re *regexp.Regexp = regexp.MustCompile(`href="/mangas/(\d*)/(.*?)">\s*([^<]+)\s*<`)
	var matches [][]string = re.FindAllStringSubmatch(page, -1)
	if len(matches) == 0 {
		return nil, errors.New("unable to parse series from page")
	}

	return []Series{}, nil
}

func getSeriesPage() (string, error) {
	resp, err := http.Get("https://onepiecechapters.com/projects")
	if err == nil {
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err == nil {
			return string(body), nil
		}
	}
	return "", err
}
