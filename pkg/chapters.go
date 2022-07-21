package pkg

import (
	"errors"
	"fmt"
	"html"
	"regexp"
	"strings"
)

// Identification information for a chapter of a mange series on TCBScans.
type Chapter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// GetChapters returns a list of available chapters on TCBScans for a series.
func GetChapters(series Series) ([]Chapter, error) {
	var url string = fmt.Sprintf("https://onepiecechapters.com/mangas/%s/%s", series.Id, series.Slug)
	page, err := getPage(url)
	if err != nil {
		return nil, err
	}

	var re *regexp.Regexp = regexp.MustCompile(`(?s)href="/chapters/(\d*)/([^"]*)"[^>]*>(.*?)</a>`)
	var matches [][]string = re.FindAllStringSubmatch(page, -1)
	if len(matches) == 0 {
		return nil, errors.New("unable to parse chapters from page")
	}

	var chapters []Chapter = make([]Chapter, 0)
	for _, match := range matches {
		name, err := getChapterName(match[3])
		if err != nil {
			continue
		}

		chapters = append(chapters, Chapter{
			Id:   match[1],
			Name: name,
			Slug: match[2],
		})
	}

	return chapters, nil
}

// getChapterName extracts the name of the chapter from the div html element
func getChapterName(div string) (string, error) {
	var name string = strings.TrimSpace(div)
	if name == "" {
		// chapter regex is greedy and matches a empty placeholder div
		return "", errors.New("empty chapter name, skip div")
	}

	var re *regexp.Regexp = regexp.MustCompile(`<div[^>]*>\s*([^<]+)\s*</div>`)
	var matches [][]string = re.FindAllStringSubmatch(name, -1)
	if len(matches) == 0 {
		return "", errors.New("unable to parse chapter name from div")
	}

	// reduce [][]string to []string
	var names []string = make([]string, 0)
	for _, match := range matches {
		names = append(names, match[1])
	}

	name = strings.Join(names[:], ": ")
	name = html.UnescapeString(name)
	return name, nil
}
