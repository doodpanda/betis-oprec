package libraryHandlers

import (
	"github.com/google/uuid"

	"betis-oprec/internal/model"
)

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
	MagicType BookType  `json:"magic_type"`
}

type BookCreateRequest struct {
	Title     string   `json:"title"`
	MagicType BookType `json:"magic_type"`
	Status    bool     `json:"status"`
}

type BookUpdateRequest struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	MagicType BookType `json:"magic_type"`
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
	Books     []model.MagicBook `json:"magic_books"`
	Total     int               `json:"total"`
	HasNext   bool              `json:"has_next"`
	Page      int               `json:"page"`
	TotalPage int               `json:"total_page"`
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
