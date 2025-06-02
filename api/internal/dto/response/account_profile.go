package response

import (
	"lv99/internal/model"
	"time"
)

// ============================
// AccountProfile
// ============================

type AccountProfile struct {
	AccountId   int       `json:"account_id"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
	AvatarURL   string    `json:"avatar_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromModelAccountProfile(m model.AccountProfile) AccountProfile {
	return AccountProfile{
		AccountId:   m.AccountId,
		DisplayName: m.DisplayName,
		Bio:         m.Bio,
		AvatarURL:   m.AvatarURL,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func FromModelAccountProfileList(models []model.AccountProfile) []AccountProfile {
	res := make([]AccountProfile, 0, len(models))
	for _, m := range models {
		res = append(res, FromModelAccountProfile(m))
	}
	return res
}
