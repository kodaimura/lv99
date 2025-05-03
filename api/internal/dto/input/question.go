package input

type Question struct {
	QuestionId      int
	QuestionTitle   string
	QuestionContent string
	QuestionAnswer  string
	QuestionLevel   int
}

type QuestionPK struct {
	QuestionId int
}