package answer

type GetDto struct {
	AccountId  int
	QuestionId int
}

type GetOneDto struct {
	Id        int
	AccountId int
}

type CreateOneDto struct {
	QuestionId int
	AccountId  int
	CodeDef    string
	CodeCall   string
}

type UpdateOneDto struct {
	Id        int
	AccountId int
	CodeDef   string
	CodeCall  string
}

type DeleteOneDto struct {
	Id        int
	AccountId int
}
