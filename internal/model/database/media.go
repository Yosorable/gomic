package database

type Author struct {
	Base
	Name        string `gorm:"uniqueIndex"`
	CoverFileID *uint
}

type Archive struct {
	Base
	Name        string `gorm:"index"`
	CoverFileID *uint
	AuthorID    *uint `gorm:"index"`
}

type ArchiveFileType uint8

const (
	PICTURE_FILE ArchiveFileType = iota
	VIDEO_FILE
)

type ArchiveFile struct {
	Base
	Name      string
	Path      string
	ArchiveID uint `gorm:"index"`
	FileType  ArchiveFileType
}

type CacheFile struct {
	Base
	Path string
}
