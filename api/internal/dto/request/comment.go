package request

type CommentUri struct {
	CommentId int `uri:"comment_id" binding:"required"`
}

type CommentBody struct {
	Content string `json:"content" binding:"required"`
}

type PostComment struct {
	AnswerUri
	CommentBody
}

type PutComment struct {
	CommentUri
	CommentBody
}