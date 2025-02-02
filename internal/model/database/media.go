package database

type Author struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Folder      string `json:"folder"`
	CoverFileId int    `json:"cover_file_id"`
}

type Archive struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CoverFileId int    `json:"cover_file_id"`
	AuthorID    int    `json:"author_id"`
}

type ArchiveFileType uint8

const (
	PICTURE_FILE ArchiveFileType = iota
	VIDEO_FILE
)

type ArchiveFile struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Path      string          `json:"path"`
	ArchiveID int             `json:"archive_id"`
	FileType  ArchiveFileType `json:"file_type"`
}
