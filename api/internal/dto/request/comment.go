package request

type Comment struct {
	CommentContent string `json:"comment_content"`
}

type CommentPK struct {
	CommentId int `uri:"answer_id"`
}

type PutComment struct {
	AnswerId  int `uri:"answer_id"`
	CommentId int `uri:"comment_id"`
}

type DeleteComment struct {
	AnswerId  int `uri:"answer_id"`
	CommentId int `uri:"comment_id"`
}
