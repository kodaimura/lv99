package input

type Comment struct {
	CommentId      int
	AnswerId       int
	AccountId      int
	CommentContent string
}

type CommentPK struct {
	CommentId int
}
