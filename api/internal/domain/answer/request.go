package answer

type QuestionUri struct {
	QuestionId int `uri:"question_id" binding:"required"`
}

type AnswerUri struct {
	AnswerId int `uri:"answer_id" binding:"required"`
}

type QuestionAnswerUri struct {
	QuestionId int `uri:"question_id" binding:"required"`
	AnswerId   int `uri:"answer_id" binding:"required"`
}

type PostOneRequest struct {
	CodeDef  string `json:"code_def" binding:"required"`
	CodeCall string `json:"code_call" binding:"required"`
}

type PutOneRequest struct {
	CodeDef  string `json:"code_def" binding:"required"`
	CodeCall string `json:"code_call" binding:"required"`
}
