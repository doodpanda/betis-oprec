package witchHandlers

import (
	"betis-oprec/internal/model"

	"github.com/google/uuid"
)

type WitchInfoResponse struct {
	ID       uuid.UUID                `json:"Id"`
	Name     string                   `json:"Name"`
	Rank     string                   `json:"Rank"`
	Accesses []model.AccessPermission `json:"AccessPermission"`
}

type WitchCreateRequest struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
	Role string `json:"Role"`
	Rank string `json:"Rank"`
}

type WitchUpdateRequest struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
	Age  int    `json:"Age"`
	Rank string `json:"Rank"`
}

type WitchDeleteRequest struct {
	ID string `json:"Id"`
}

type WitchListRequest struct {
	Page  int `json:"Page"`
	Limit int `json:"Limit"`
}

type WitchListResponse struct {
	Witches []WitchInfoResponse `json:"Witches"`
	Total   int                 `json:"Total"`
	HasNext bool                `json:"HasNext"`
}
