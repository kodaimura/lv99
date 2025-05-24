package input

import "time"

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

type GetChat struct {
	FromId  int
	ToId    int
	Before  time.Time
	Limit   int
}
