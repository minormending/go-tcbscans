package pkg

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func GetSeries() (string, error) {
	resp, err := http.Get("https://onepiecechapters.com/projects")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	page := string(body)
	re := regexp.MustCompile(`href="/mangas/(\d*)/(.*?)">\s*([^<]+)\s*<`)
	matches := re.FindAllStringSubmatch(page, -1)
	for i := range matches {
		fmt.Printf("%s %s %s\n", matches[i][1], matches[i][2], matches[i][3])
	}

	return "", nil
}
