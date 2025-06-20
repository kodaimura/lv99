package comment

type AnswerUri struct {
	AnswerId int `uri:"answer_id" binding:"required"`
}

type CommentUri struct {
	CommentId int `uri:"comment_id" binding:"required"`
}

type PostOneRequest struct {
	Content string `json:"content" binding:"required"`
}

type PutOneRequest struct {
	Content string `json:"content" binding:"required"`
}
