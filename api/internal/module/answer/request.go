package answer

type AnswerUri struct {
	AnswerId int `uri:"answer_id" binding:"required"`
}

type QuestionAnswerUri struct {
	QuestionId int `uri:"question_id" binding:"required"`
	AnswerId   int `uri:"answer_id" binding:"required"`
}

type GetRequest struct {
	QuestionId int `form:"question_id"`
}

type PostOneRequest struct {
	QuestionId int    `json:"question_id" binding:"required"`
	CodeDef    string `json:"code_def" binding:"required"`
	CodeCall   string `json:"code_call" binding:"required"`
}

type PutOneRequest struct {
	CodeDef  string `json:"code_def" binding:"required"`
	CodeCall string `json:"code_call" binding:"required"`
}

type AdminGetRequest struct {
	AccountId  int `form:"account_id"`
	QuestionId int `form:"question_id"`
}
