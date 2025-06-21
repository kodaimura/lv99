package auth

// ============================
// Login
// ============================

type LoginResponse struct {
	AccountId        int    `json:"account_id"`
	AccountRole      int    `json:"account_role"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	AccessExpiresIn  int    `json:"access_expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

// ============================
// Refresh
// ============================

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
