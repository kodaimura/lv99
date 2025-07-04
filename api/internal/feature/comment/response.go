package comment

import "time"

// ============================
// CommentWithProfileResponse
// ============================

type CommentWithProfileResponse struct {
	Id          int       `json:"id"`
	AnswerId    int       `json:"answer_id"`
	AccountId   int       `json:"account_id"`
	DisplayName string    `json:"display_name"`
	AvatarURL   string    `json:"avatar_url"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToCommentWithProfileResponse(m CommentWithProfile) CommentWithProfileResponse {
	return CommentWithProfileResponse(m)
}

func ToCommentWithProfileResponseList(models []CommentWithProfile) []CommentWithProfileResponse {
	res := make([]CommentWithProfileResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToCommentWithProfileResponse(m))
	}
	return res
}

// ============================
// CommentCountResponse
// ============================

type CommentCountResponse struct {
	QuestionId    int       `json:"question_id"`
	QuestionTitle string    `json:"question_title"`
	QuestionLevel int       `json:"question_level"`
	AnswerId      int       `json:"answer_id"`
	CommentCount  int       `json:"comment_count"`
	CreatedAt     time.Time `json:"created_at"`
}

func ToCommentCountResponse(m CommentCount) CommentCountResponse {
	return CommentCountResponse(m)
}

func ToCommentCountResponseList(models []CommentCount) []CommentCountResponse {
	res := make([]CommentCountResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToCommentCountResponse(m))
	}
	return res
}
