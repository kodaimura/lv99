package answer

type SearchRequest struct {
	QuestionId       int   `form:"question_id"`
	AccountId        int   `form:"account_id"`
	Level            int   `form:"level"`
	IsCorrect        *bool `form:"is_correct"`
	CommentAccountId int   `form:"comment_account_id"`
}

type GetStatusRequest struct {
	AccountId int `form:"account_id"`
}
