package pkg

import (
	"errors"
	"fmt"
	"regexp"
)

func GetChapterPages(chapter Chapter) ([]string, error) {
	var url string = fmt.Sprintf("https://onepiecechapters.com/chapters/%s/%s", chapter.Id, chapter.Slug)
	page, err := getPage(url)
	if err != nil {
		return nil, err
	}

	var re *regexp.Regexp = regexp.MustCompile(`(?s)<picture[^>]*>.*?<img[^>]*?src="([^"]+)"[^>]*>.*?</picture>`)
	var matches [][]string = re.FindAllStringSubmatch(page, -1)
	if len(matches) == 0 {
		return nil, errors.New("unable to parse chapter images from page")
	}

	var images []string = make([]string, 0)
	for _, match := range matches {
		images = append(images, match[1])
	}

	return images, nil
}
