package account_extended

import (
	usecase "lv99/internal/usecase/account_extended"
	"time"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type AccountWithProfileResponse struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Role        int        `json:"role"`
	DisplayName string     `json:"display_name"`
	Bio         string     `json:"bio"`
	AvatarURL   string     `json:"avatar_url"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func ToAccountWithProfileResponse(m usecase.AccountWithProfile) AccountWithProfileResponse {
	return AccountWithProfileResponse(m)
}

func ToAccountWithProfileResponseList(models []usecase.AccountWithProfile) []AccountWithProfileResponse {
	res := make([]AccountWithProfileResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAccountWithProfileResponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type AccountUri struct {
	AccountId int `uri:"account_id" binding:"required"`
}
