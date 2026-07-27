package chat

import "time"

type GetDto struct {
	FromId      int
	ToId        int
	AccountRole int
}

type CreateOneDto struct {
	FromId      int
	ToId        int
	AccountRole int
	Message     string
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

type ReadDto struct {
	ToId        int
	FromId      int
	AccountRole int
}

type PaginateDto struct {
	FromId      int
	ToId        int
	AccountRole int
	Before      time.Time
	Limit       int
}
