package comment

type CommentUri struct {
	CommentId int `uri:"comment_id" binding:"required"`
}

type GetRequest struct {
	AnswerId int `form:"answer_id"`
}

type PostOneRequest struct {
	AnswerId int    `json:"answer_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type PutOneRequest struct {
	Content string `json:"content" binding:"required"`
}
