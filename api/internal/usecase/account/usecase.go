package account

import (
	accountModule "lv99/internal/module/account"

	"gorm.io/gorm"
)

type Usecase interface {
	Get(in GetDto, db *gorm.DB) ([]accountModule.Account, error)
	GetOne(in GetOneDto, db *gorm.DB) (accountModule.Account, error)
	UpdateOne(in UpdateOneDto, db *gorm.DB) (accountModule.Account, error)
	DeleteOne(in DeleteOneDto, db *gorm.DB) error
}

type usecase struct {
	accountService accountModule.Service
}

func NewUsecase(accountService accountModule.Service) Usecase {
	return &usecase{
		accountService: accountService,
	}
}

func (uc *usecase) Get(in GetDto, db *gorm.DB) ([]accountModule.Account, error) {
	return uc.accountService.Get(accountModule.Account{}, db)
}

func (uc *usecase) GetOne(in GetOneDto, db *gorm.DB) (accountModule.Account, error) {
	return uc.accountService.GetOne(accountModule.Account{Id: in.Id}, db)
}

func (uc *usecase) UpdateOne(in UpdateOneDto, db *gorm.DB) (accountModule.Account, error) {
	return uc.accountService.UpdateOne(accountModule.Account{Id: in.Id, Name: in.Name}, db)
}

func (uc *usecase) DeleteOne(in DeleteOneDto, db *gorm.DB) error {
	return uc.accountService.DeleteOne(accountModule.Account{Id: in.Id}, db)
}
