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
	ID        uuid.UUID `json:"Id"`
	Title     string    `json:"Title"`
	MagicType BookType  `json:"MagicType"`
}

type BookCreateRequest struct {
	Title     string   `json:"Title"`
	MagicType BookType `json:"MagicType"`
	Status    bool     `json:"IsAvailable"`
}

type BookUpdateRequest struct {
	ID        string   `json:"Id"`
	Title     string   `json:"Title"`
	MagicType BookType `json:"MagicType"`
	Status    bool     `json:"IsAvailable"`
}

type BookDeleteRequest struct {
	ID string `json:"Id"`
}

type BookListRequest struct {
	Page  int `json:"Page"`
	Limit int `json:"Limit"`
}

type BookListResponse struct {
	Books     []model.MagicBook `json:"MagicBooks"`
	Total     int               `json:"Total"`
	HasNext   bool              `json:"HasNext"`
	Page      int               `json:"Page"`
	TotalPage int               `json:"TotalPage"`
}

type BookDetailRequest struct {
	ID string `json:"Id"`
}

type BookDetailResponse struct {
	ID         string `json:"Id"`
	Title      string `json:"Title"`
	MagicType  string `json:"Type"`
	CreateDate string `json:"CreateDate"`
	Status     bool   `json:"IsAvailable"`
}
