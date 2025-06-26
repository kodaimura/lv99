package answer_search

type SearchDto struct {
	QuestionId       int   `json:"question_id"`
	AccountId        int   `json:"account_id"`
	IsCorrect        *bool `json:"is_correct"`
	CommentAccountId int   `json:"comment_account_id"`
}
