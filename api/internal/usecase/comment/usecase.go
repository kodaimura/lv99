package comment

import (
	commentModule "lv99/internal/module/comment"

	"gorm.io/gorm"
)

type Usecase interface {
	Get(in GetDto) ([]commentModule.Comment, error)
	GetOne(in GetOneDto) (commentModule.Comment, error)
	CreateOne(in CreateOneDto) (commentModule.Comment, error)
	UpdateOne(in UpdateOneDto) (commentModule.Comment, error)
	DeleteOne(in DeleteOneDto) error
}

type usecase struct {
	db             *gorm.DB
	commentService commentModule.Service
}

func NewUsecase(
	db *gorm.DB,
	commentService commentModule.Service,
) Usecase {
	return &usecase{
		db:             db,
		commentService: commentService,
	}
}

func (uc *usecase) Get(in GetDto) ([]commentModule.Comment, error) {
	return uc.commentService.Get(commentModule.Comment{
		AnswerId:  in.AnswerId,
		AccountId: in.AccountId,
	}, uc.db)
}

func (uc *usecase) GetOne(in GetOneDto) (commentModule.Comment, error) {
	return uc.commentService.GetOne(commentModule.Comment{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, uc.db)
}

func (uc *usecase) CreateOne(in CreateOneDto) (commentModule.Comment, error) {
	return uc.commentService.CreateOne(commentModule.Comment{
		AnswerId:  in.AnswerId,
		AccountId: in.AccountId,
		Content:   in.Content,
	}, uc.db)
}

func (uc *usecase) UpdateOne(in UpdateOneDto) (commentModule.Comment, error) {
	return uc.commentService.CreateOne(commentModule.Comment{
		Id:        in.Id,
		AccountId: in.AccountId,
		Content:   in.Content,
	}, uc.db)
}

func (uc *usecase) DeleteOne(in DeleteOneDto) error {
	return uc.commentService.DeleteOne(commentModule.Comment{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, uc.db)
}
