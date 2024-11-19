package witchHandlers

import "github.com/google/uuid"

type WitchInfoResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Rank string    `json:"rank"`
}

type WitchCreateRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role"`
	Rank string `json:"rank"`
}

type WitchUpdateRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Rank string `json:"rank"`
}

type WitchDeleteRequest struct {
	ID string `json:"id"`
}

type WitchListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type WitchListResponse struct {
	Witches []WitchInfoResponse `json:"witches"`
	Total   int                 `json:"total"`
	HasNext bool                `json:"has_next"`
}
