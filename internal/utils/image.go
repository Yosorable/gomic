package utils

import (
	"os"

	"github.com/disintegration/imaging"
)

func CreateImageThumb(filePath, thumbFilePath string) error {
	if FileExists(thumbFilePath) {
		return nil
	}
	fileImg, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fileImg.Close()

	src, err := imaging.Open(filePath)
	if err != nil {
		return err
	}

	dsc := imaging.Fill(src, 128, 192, imaging.Center, imaging.Lanczos)
	return imaging.Save(dsc, thumbFilePath)
}
