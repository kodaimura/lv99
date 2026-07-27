package comment

type GetDto struct {
	AnswerId    int
	AccountId   int
	AccountRole int
}

type GetOneDto struct {
	Id        int
	AccountId int
}

type CreateOneDto struct {
	AnswerId    int
	AccountId   int
	AccountRole int
	Content     string
}

type UpdateOneDto struct {
	Id        int
	AccountId int
	Content   string
}

type DeleteOneDto struct {
	Id        int
	AccountId int
}
