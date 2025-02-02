package global

import (
	"sync"

	"github.com/Yosorable/gomic/internal/model"
)

var (
	CONFIG Config
)

var (
	SCANNING_MUTEX      = &sync.Mutex{}
	IS_SERVER_SCANNING  = false
	MEDIA_SERVER_PREFIX = "/media"
)

var TMP_MEMORY_DB []*model.AuthorMedia = []*model.AuthorMedia{}
var TMP_SYNC_RECORD []string = []string{}
