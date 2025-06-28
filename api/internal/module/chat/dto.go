package chat

import "time"

type GetDto struct {
	FromId int
	ToId   int
	Before time.Time
	Limit  int
}

type CreateOneDto struct {
	FromId  int
	ToId    int
	Message string
}

type UpdateOneDto struct {
	Id      int
	FromId  int
	ToId    int
	Message string
	IsRead  bool
}

type DeleteOneDto struct {
	Id int
}
