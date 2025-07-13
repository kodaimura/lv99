package answer

import (
	"gorm.io/gorm"
)

type Service interface {
	Get(in Answer, db *gorm.DB) ([]Answer, error)
	GetOne(in Answer, db *gorm.DB) (Answer, error)
	CreateOne(in Answer, db *gorm.DB) (Answer, error)
	UpdateOne(in Answer, db *gorm.DB) (Answer, error)
	DeleteOne(in Answer, db *gorm.DB) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (srv *service) Get(in Answer, db *gorm.DB) ([]Answer, error) {
	return srv.repository.Get(&Answer{
		QuestionId: in.QuestionId,
		AccountId:  in.AccountId,
	}, db)
}

func (srv *service) GetOne(in Answer, db *gorm.DB) (Answer, error) {
	return srv.repository.GetOne(&Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
}

func (srv *service) CreateOne(in Answer, db *gorm.DB) (Answer, error) {
	return srv.repository.Insert(&in, db)
}

func (srv *service) UpdateOne(in Answer, db *gorm.DB) (Answer, error) {
	ans, err := srv.repository.GetOne(&Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
	if err != nil {
		return Answer{}, err
	}

	ans.CodeDef = in.CodeDef
	ans.CodeCall = in.CodeCall
	ans.CallOutput = in.CallOutput
	ans.CallError = in.CallError
	ans.IsCorrect = in.IsCorrect

	return srv.repository.Update(&ans, db)
}

func (srv *service) DeleteOne(in Answer, db *gorm.DB) error {
	return srv.repository.Delete(&Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
}
