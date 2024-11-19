package libraryHandlers

import "github.com/google/uuid"

type BookType string

const (
	Elemental  BookType = "elemental"
	Illusion   BookType = "illusion"
	Necromancy BookType = "necromancy"
	Healing    BookType = "healing"
)

type BookInfoResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	MagicType BookType  `json:"type"`
}

type BookCreateRequest struct {
	Title     string   `json:"title"`
	MagicType BookType `json:"type"`
	Status    bool     `json:"status"`
}

type BookUpdateRequest struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	MagicType BookType `json:"type"`
	Status    bool     `json:"status"`
}

type BookDeleteRequest struct {
	ID string `json:"id"`
}

type BookListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type BookListResponse struct {
	Books   []BookInfoResponse `json:"books"`
	Total   int                `json:"total"`
	HasNext bool               `json:"has_next"`
}

type BookDetailRequest struct {
	ID string `json:"id"`
}

type BookDetailResponse struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	MagicType  string `json:"type"`
	CreateDate string `json:"create_date"`
	Status     bool   `json:"status"`
}
