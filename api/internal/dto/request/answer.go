package request

type Answer struct {
	CodeDef  string `json:"code_def"`
	CodeCall string `json:"code_call"`
}

type AnswerPK struct {
	AnswerId int `uri:"answer_id"`
}

type PutAnswer struct {
	QuestionId int `uri:"question_id"`
	AnswerId   int `uri:"answer_id"`
}

type DeleteAnswer struct {
	QuestionId int `uri:"question_id"`
	AnswerId   int `uri:"answer_id"`
}
