package request

type Question struct {
	QuestionTitle   string `json:"question_title"`
	QuestionContent string `json:"question_content"`
	QuestionAnswer  string `json:"question_answer"`
	QuestionLevel   int    `json:"question_level"`
}

type QuestionPK struct {
	QuestionId int `uri:"question_id"`
}
