package request

type CommentPK struct {
	Id int `uri:"id"`
}

type GetComment struct {
	AnswerId  int `uri:"answer_id"`
}

type PostComment struct {
	AnswerId  int `uri:"answer_id"`
	Content string `json:"content"`
}

type PutComment struct {
	Id int `uri:"id"`
	Content string `json:"content"`
}