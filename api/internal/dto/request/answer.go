package request

type Answer struct {
	CodeDef  string `json:"code_def"`
	CodeCall string `json:"code_call"`
}

type AnswerPK struct {
	AnswerId int `uri:"answer_id"`
}
