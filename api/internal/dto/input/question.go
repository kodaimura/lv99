package input

type Question struct {
	Id      int
	Title   string
	Content string
	Answer  string
	Level   int
}

type QuestionPK struct {
	Id int
}
