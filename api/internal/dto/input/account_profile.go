package input

type AccountProfile struct {
	AccountId   int
	DisplayName string
	Bio         string
	AvatarURL   string
}

type AccountProfilePK struct {
	AccountId int
}
