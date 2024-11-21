package accessHandlers

import (
	"betis-oprec/internal/model"
	"time"

	"github.com/google/uuid"
)

type AccessCreateRequest struct {
	WitchID     string `json:"WitchId"`
	MagicBookID string `json:"MagicBookId"`
}

type AccessInfoResponse struct {
	ID          uuid.UUID       `json:"Id"`
	WitchID     uuid.UUID       `json:"WitchId"`
	MagicBookID uuid.UUID       `json:"BookId"`
	PermitDate  time.Time       `json:"PermitDate"`
	Witch       model.Witch     `json:"Witch"`
	MagicBook   model.MagicBook `json:"MagicBook"`
}

type AccessDeleteRequest struct {
	ID string `json:"Id"`
}

type AccessDeleteResponse struct {
	ID string `json:"Id"`
}

type AccessListRequest struct {
	Page  int `json:"Page"`
	Limit int `json:"Limit"`
}

type AccessListResponse struct {
	Accesses  []model.AccessPermission `json:"Accesses"`
	Total     int                      `json:"Total"`
	HasNext   bool                     `json:"HasNext"`
	Page      int                      `json:"Page"`
	TotalPage int                      `json:"TotalPage"`
}
