package request

type QuestionPK struct {
	Id int `uri:"id"`
}

type QuestionBody struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
    Answer  string `json:"answer" binding:"required"`
    Level   int    `json:"level" binding:"required,min=1"`
}

type PostQuestion struct {
    QuestionBody
}

type PutQuestion struct {
    Id int `uri:"id"`
    QuestionBody
}