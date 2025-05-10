package request

type AnswerUri struct {
	AnswerId int `uri:"answer_id" binding:"required"`
}

type AnswerBody struct {
	CodeDef  string `json:"code_def" binding:"required"`
	CodeCall string `json:"code_call" binding:"required"`
}

type PostAnswer struct {
	QuestionUri
	AnswerBody
}

type PutAnswer struct {
	AnswerUri
	AnswerBody
}