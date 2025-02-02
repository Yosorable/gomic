package utils

import (
	"os"
	"strings"
)

var pictureSuffix = []string{
	".jpg",
	".jpeg",
	".png",
	".gif",
	".bmp",
	".tiff",
	".tif",
	".webp",
}

var videoSuffix = []string{
	".mp4",
	".mov",
	".avi",
	".mkv",
	".wmv",
	".flv",
	".webm",
}

func IsPicture(name string) bool {
	name = strings.ToLower(name)
	for _, ele := range pictureSuffix {
		if strings.HasSuffix(name, ele) {
			return true
		}
	}
	return false
}

func IsVideo(name string) bool {
	name = strings.ToLower(name)
	for _, ele := range videoSuffix {
		if strings.HasSuffix(name, ele) {
			return true
		}
	}
	return false
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false // 其他错误
	}
	return true
}
