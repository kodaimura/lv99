package service

import (
	"lv99/internal/dto/input"
	"lv99/internal/model"
	"lv99/internal/repository"
)

type CommentService interface {
	Get(in input.Comment) ([]model.Comment, error)
	GetOne(in input.Comment) (model.Comment, error)
	CreateOne(in input.Comment) (model.Comment, error)
	UpdateOne(in input.Comment) (model.Comment, error)
	DeleteOne(in input.Comment) error
}

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(
	commentRepository repository.CommentRepository, 
) CommentService {
	return &commentService{
		commentRepository: commentRepository,
	}
}

func (srv *commentService) Get(in input.Comment) ([]model.Comment, error) {
	return srv.commentRepository.Get(&model.Comment{
		AnswerId: in.AnswerId,
		AccountId: in.AccountId,
	})
}

func (srv *commentService) GetOne(in input.Comment) (model.Comment, error) {
	return srv.commentRepository.GetOne(&model.Comment{
		Id: in.Id,
		AccountId: in.AccountId,
	})
}

func (srv *commentService) CreateOne(in input.Comment) (model.Comment, error) {
	return srv.commentRepository.Insert(&model.Comment{
		AnswerId: in.AnswerId,
		AccountId: in.AccountId,
		Content: in.Content,
	})
}

func (srv *commentService) UpdateOne(in input.Comment) (model.Comment, error) {
	comment, err := srv.commentRepository.GetOne(&model.Comment{
		Id: in.Id,
		AccountId: in.AccountId,
	})
	if err != nil {
		return model.Comment{}, err
	}

	comment.Content = in.Content
	return srv.commentRepository.Update(&comment)
}

func (srv *commentService) DeleteOne(in input.Comment) error {
	return srv.commentRepository.Delete(&model.Comment{
		Id: in.Id,
		AccountId: in.AccountId,
	})
}