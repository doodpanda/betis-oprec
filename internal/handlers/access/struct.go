package accessHandlers

import "github.com/google/uuid"

type AccessCreateRequest struct {
	WitchID string `json:"witch_id"`
	BookID  string `json:"book_id"`
}

type AccessInfoResponse struct {
	ID         uuid.UUID `json:"id"`
	WitchID    uuid.UUID `json:"witch_id"`
	BookID     uuid.UUID `json:"book_id"`
	PermitDate string    `json:"permit_date"`
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
	Accesses []AccessInfoResponse `json:"accesses"`
	Total    int                  `json:"total"`
	HasNext  bool                 `json:"has_next"`
}
