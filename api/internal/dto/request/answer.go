package request

type AnswerPK struct {
	Id int `uri:"id"`
}

type GetAnswer struct {
	QuestionId int `uri:"question_id"`
}

type PostAnswer struct {
	QuestionId int `uri:"question_id"`
	CodeDef  string `json:"code_def"`
	CodeCall string `json:"code_call"`
}

type PutAnswer struct {
	Id         int `uri:"id"`
	CodeDef  string `json:"code_def"`
	CodeCall string `json:"code_call"`
}