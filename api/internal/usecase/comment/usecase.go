package comment

import (
	commentModule "lv99/internal/module/comment"

	"gorm.io/gorm"
)

type Usecase interface {
	Get(in GetDto, db *gorm.DB) ([]commentModule.Comment, error)
	GetOne(in GetOneDto, db *gorm.DB) (commentModule.Comment, error)
	CreateOne(in CreateOneDto, db *gorm.DB) (commentModule.Comment, error)
	UpdateOne(in UpdateOneDto, db *gorm.DB) (commentModule.Comment, error)
	DeleteOne(in DeleteOneDto, db *gorm.DB) error
}

type usecase struct {
	commentService commentModule.Service
}

func NewUsecase(commentService commentModule.Service) Usecase {
	return &usecase{
		commentService: commentService,
	}
}

func (srv *usecase) Get(in GetDto, db *gorm.DB) ([]commentModule.Comment, error) {
	return srv.commentService.Get(commentModule.Comment{
		AnswerId:  in.AnswerId,
		AccountId: in.AccountId,
	}, db)
}

func (srv *usecase) GetOne(in GetOneDto, db *gorm.DB) (commentModule.Comment, error) {
	return srv.commentService.GetOne(commentModule.Comment{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
}

func (srv *usecase) CreateOne(in CreateOneDto, db *gorm.DB) (commentModule.Comment, error) {
	return srv.commentService.CreateOne(commentModule.Comment{
		AnswerId:  in.AnswerId,
		AccountId: in.AccountId,
		Content:   in.Content,
	}, db)
}

func (srv *usecase) UpdateOne(in UpdateOneDto, db *gorm.DB) (commentModule.Comment, error) {
	return srv.commentService.CreateOne(commentModule.Comment{
		Id:        in.Id,
		AccountId: in.AccountId,
		Content:   in.Content,
	}, db)
}

func (srv *usecase) DeleteOne(in DeleteOneDto, db *gorm.DB) error {
	return srv.commentService.DeleteOne(commentModule.Comment{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
}
