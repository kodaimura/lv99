package account_with_profile

type AccountUri struct {
	AccountId int `uri:"account_id" binding:"required"`
}
