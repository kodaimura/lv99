package input

type Admin struct {
	AdminId       int
	AdminName     string
	AdminPassword string
}

type AdminPK struct {
	AdminId int
}

type AdminSignup struct {
	AdminName     string
	AdminPassword string
}

type AdminLogin struct {
	AdminName     string
	AdminPassword string
}
