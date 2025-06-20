package account_profile

type PutMeRequest struct {
	DisplayName string `json:"display_name" binding:"required"`
	Bio         string `json:"bio" binding:"omitempty"`
	AvatarURL   string `json:"avatar_url" binding:"omitempty,url"`
}
