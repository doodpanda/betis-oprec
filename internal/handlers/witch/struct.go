package witchHandlers

import "github.com/google/uuid"

type WitchInfoResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Role string    `json:"role"`
}

type WitchCreateRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role"`
}

type WitchUpdateRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role"`
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
