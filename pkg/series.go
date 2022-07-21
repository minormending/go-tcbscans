package pkg

import (
	"errors"
	"io"
	"net/http"
	"regexp"
)

// A manga series identification information on TCBScans.
type Series struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// GetSeries returns a slice of all the series found on the TCBScans website.
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

	var series []Series = make([]Series, 0)
	for _, match := range matches {
		series = append(series, Series{
			Id:   match[1],
			Name: match[2],
			Slug: match[3],
		})
	}

	return series, nil
}

// getSeriesPage fetches the html page from TCBScans containing the list of series.
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
