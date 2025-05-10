package request

type QuestionUri struct {
	QuestionId int `uri:"question_id" binding:"required"`
}

type QuestionBody struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Level   int    `json:"level" binding:"required,min=1"`
}
