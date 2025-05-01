package request

type AdminSignup struct {
	AdminName     string `json:"admin_name" binding:"required"`
	AdminPassword string `json:"admin_password" binding:"required,min=8"`
}

type AdminLogin struct {
	AdminName     string `json:"admin_name"`
	AdminPassword string `json:"admin_password"`
}

type PutAdmin struct {
	AdminName string `json:"admin_name" binding:"required"`
}

type PutAdminPassword struct {
	OldAdminPassword string `json:"old_admin_password" binding:"required"`
	NewAdminPassword string `json:"new_admin_password" binding:"required"`
}
