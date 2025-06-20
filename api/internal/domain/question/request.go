package question

type QuestionUri struct {
	QuestionId int `uri:"question_id" binding:"required"`
}

type PostOneRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Level   int    `json:"level" binding:"required,min=1"`
}

type PutOneRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Level   int    `json:"level" binding:"required,min=1"`
}
