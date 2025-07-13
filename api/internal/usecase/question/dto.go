package question

type GetDto struct {
	Title   string
	Content string
	Answer  string
	Level   int
}

type GetAllDto struct {
	Title   string
	Content string
	Answer  string
	Level   int
}

type GetOneDto struct {
	Id int
}

type CreateOneDto struct {
	Title   string
	Content string
	Answer  string
	Level   int
}

type UpdateOneDto struct {
	Id      int
	Title   string
	Content string
	Answer  string
	Level   int
}

type DeleteOneDto struct {
	Id int
}

type RestoreOneDto struct {
	Id int
}
