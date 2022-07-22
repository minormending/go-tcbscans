package pkg

import (
	"errors"
	"fmt"
	"image/png"
	"log"
	"os"
	"path"
)

// SaveChapter downloads all the manga images for the given chapter into the given directory.
func SaveChapter(chapter Chapter, directory string, force bool) error {
	var folder string = path.Join(directory, chapter.Slug)
	if _, err := os.Stat(folder); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return err
		}
	} else if !force {
		log.Printf("the chapter %s already exists in the directory %s, skipping it", chapter.Slug, directory)
		//return nil
	}

	images, err := GetChapterPages(chapter)
	if err != nil {
		return err
	}

	for idx, url := range images {
		var ext string = path.Ext(url)
		var filename string = fmt.Sprintf("%s/%s-page%d%s", folder, chapter.Slug, idx+1, ext)

		err = savePage(url, filename)
		if err != nil {
			return err
		}
	}
	return nil
}

