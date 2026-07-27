package comment

import (
	"lv99/internal/core"
	"lv99/internal/helper"
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
	if in.AccountRole != helper.ACCOUNT_ROLE_ADMIN {
		var count int64
		err := uc.db.Table("answer").
			Where("id = ? AND account_id = ? AND deleted_at IS NULL", in.AnswerId, in.AccountId).
			Count(&count).Error
		if err != nil {
			return nil, core.NewUnexpectedError(err)
		}
		if count == 0 {
			return nil, core.ErrForbidden
		}
	}

	return uc.commentService.Get(commentModule.Comment{
		AnswerId: in.AnswerId,
	}, uc.db)
}

func (uc *usecase) GetOne(in GetOneDto) (commentModule.Comment, error) {
	return uc.commentService.GetOne(commentModule.Comment{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, uc.db)
}

func (uc *usecase) CreateOne(in CreateOneDto) (commentModule.Comment, error) {
	if in.AccountRole != helper.ACCOUNT_ROLE_ADMIN {
		var count int64
		err := uc.db.Table("answer").
			Where("id = ? AND account_id = ? AND deleted_at IS NULL", in.AnswerId, in.AccountId).
			Count(&count).Error
		if err != nil {
			return commentModule.Comment{}, core.NewUnexpectedError(err)
		}
		if count == 0 {
			return commentModule.Comment{}, core.ErrForbidden
		}
	}

	return uc.commentService.CreateOne(commentModule.Comment{
		AnswerId:  in.AnswerId,
		AccountId: in.AccountId,
		Content:   in.Content,
	}, uc.db)
}

func (uc *usecase) UpdateOne(in UpdateOneDto) (commentModule.Comment, error) {
	return uc.commentService.UpdateOne(commentModule.Comment{
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
