package account

type AccountUri struct {
	AccountId int `uri:"account_id" binding:"required"`
}

type PutMeRequest struct {
	Name string `json:"name" binding:"required"`
}
