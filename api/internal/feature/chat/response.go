package chat

import "time"

// ============================
// UnreadCountResponse
// ============================

type UnreadCountResponse struct {
	AccountId   int       `json:"account_id"`
	UnreadCount int       `json:"unread_count"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToUnreadCountResponse(m UnreadCount) UnreadCountResponse {
	return UnreadCountResponse(m)
}

func ToUnreadCountResponseList(models []UnreadCount) []UnreadCountResponse {
	res := make([]UnreadCountResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToUnreadCountResponse(m))
	}
	return res
}
