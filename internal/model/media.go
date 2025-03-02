package model

type ArchiveFile struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Path string `json:"path"`
	URL  string `json:"url"`
}

type Archive struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	Files    []ArchiveFile `json:"files,omitempty"`
	CoverURL string        `json:"cover_url"`
}

type AuthorMedia struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Archives []Archive `json:"archives,omitempty"`
	CoverURL string    `json:"cover_url"`
}
