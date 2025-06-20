package account

import (
	"time"
)

// ============================
// Account
// ============================

type AccountResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToAccountResponse(m Account) AccountResponse {
	return AccountResponse{
		Id:        m.Id,
		Name:      m.Name,
		Role:      m.Role,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToAccountResponseList(models []Account) []AccountResponse {
	res := make([]AccountResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAccountResponse(m))
	}
	return res
}
