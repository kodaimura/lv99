package account_profile

import (
	"time"
)

// ============================
// AccountProfileResponse
// ============================

type AccountProfileResponse struct {
	AccountId   int       `json:"account_id"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
	AvatarURL   string    `json:"avatar_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToAccountProfileResponse(m AccountProfile) AccountProfileResponse {
	return AccountProfileResponse{
		AccountId:   m.AccountId,
		DisplayName: m.DisplayName,
		Bio:         m.Bio,
		AvatarURL:   m.AvatarURL,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func ToAccountProfileResponseList(models []AccountProfile) []AccountProfileResponse {
	res := make([]AccountProfileResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAccountProfileResponse(m))
	}
	return res
}
