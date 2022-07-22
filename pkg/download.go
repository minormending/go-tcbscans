package pkg

import (
	"errors"
	"fmt"
	"image/png"
	"log"
	"os"
	"path"

	"github.com/yusukebe/go-pngquant"
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

		err = minimizePng(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func minimizePng(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	var imgSizeBefore int64 = 0
	imgInfo, err := file.Stat()
	if err != nil {
		imgSizeBefore = imgInfo.Size()
	}

	img, err := png.Decode(file)
	file.Close()
	if err != nil {
		return err
	}

	minimizedImg, err := pngquant.Compress(img, "3")
	if err != nil {
		return err
	}

	file, err = os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, minimizedImg)
	if err != nil {
		return err
	}

	if imgSizeBefore > 0 {
		imgInfo, err := file.Stat()
		if err == nil {
			imgSizeAfter := imgInfo.Size()
			var change int64 = int64((float64(imgSizeAfter) / float64(imgSizeBefore)) * 100)
			log.Printf("compressed image %s by %d%%", filename, change)
		}
	}

	return nil
}
