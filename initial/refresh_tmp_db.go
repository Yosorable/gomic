package initial

import "github.com/Yosorable/gomic/internal/handler"

func RefreshTmpDB() {
	handler.ScanDirs()
}
