package input

type Chat struct {
	Id      int
	FromId  int
	ToId    int
	Message string
	IsRead  bool
}

type ChatPK struct {
	Id int
}
