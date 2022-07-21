package pkg

import (
	"fmt"
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
		return []Series{}, err
	}

	var re *regexp.Regexp = regexp.MustCompile(`href="/mangas/(\d*)/(.*?)">\s*([^<]+)\s*<`)
	var matches [][]string = re.FindAllStringSubmatch(page, -1)
	for i := range matches {
		fmt.Printf("%s %s %s\n", matches[i][1], matches[i][2], matches[i][3])
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
