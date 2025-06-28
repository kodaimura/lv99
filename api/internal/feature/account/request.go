package account

type AccountUri struct {
	AccountId int `uri:"account_id" binding:"required"`
}
