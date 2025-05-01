package response

type Admin struct {
	AdminId   int    `json:"admin_id"`
	AdminName string `json:"admin_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AdminLogin struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	AccessExpiresIn  int    `json:"access_expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	Admin            Admin  `json:"admin"`
}

type AdminRefresh struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
