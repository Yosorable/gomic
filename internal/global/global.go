package global

import (
	"sync"

	"github.com/Yosorable/gomic/internal/config"
	"gorm.io/gorm"
)

var (
	CONFIG config.Config
)

var (
	SCANNING_MUTEX      = &sync.Mutex{}
	IS_SERVER_SCANNING  = false
	MEDIA_SERVER_PREFIX = "/media"

	DB             *gorm.DB
	SKIP_AUTH_PATH = []string{"/login", "/media", "/thumb"}
)

var TMP_SYNC_RECORD []string = []string{}
