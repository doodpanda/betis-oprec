package accessHandlers

import (
	"betis-oprec/internal/model"

	"github.com/google/uuid"
)

type AccessCreateRequest struct {
	WitchID     string `json:"witch_id"`
	MagicBookID string `json:"magicbook_id"`
}

type AccessInfoResponse struct {
	ID          uuid.UUID `json:"id"`
	WitchID     uuid.UUID `json:"witch_id"`
	MagicBookID uuid.UUID `json:"book_id"`
	AccessDate  string    `json:"_date"`
}

type AccessDeleteRequest struct {
	ID string `json:"id"`
}

type AccessDeleteResponse struct {
	ID string `json:"id"`
}

type AccessListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type AccessListResponse struct {
	Accesses  []model.AccessPermission `json:"accesses"`
	Total     int                      `json:"total"`
	HasNext   bool                     `json:"has_next"`
	Page      int                      `json:"page"`
	TotalPage int                      `json:"total_page"`
}
