package input

type Answer struct {
	AnswerId   int
	QuestionId int
	AccountId  int
	CodeDef    string
	CodeCall   string
}

type AnswerPK struct {
	AnswerId int
}
