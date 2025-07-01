package comment

import "time"

type GetWithProfileRequest struct {
	AnswerId int `form:"answer_id"`
}

type GetCountRequest struct {
	Since *time.Time `form:"since" time_format:"2006-01-02"`
}
