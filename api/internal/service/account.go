package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"lv99/internal/core"
	"lv99/internal/dto/input"
	"lv99/internal/helper"
	"lv99/internal/model"
	"lv99/internal/repository"
)

type AccountService interface {
	Get(in input.Account) ([]model.Account, error)
	GetOne(in input.Account) (model.Account, error)
	CreateOne(in input.Account) (model.Account, error)
	UpdateOne(in input.Account) (model.Account, error)
	DeleteOne(in input.Account) error

	Login(in input.Login) (model.Account, error)
	UpdatePassword(in input.UpdatePassword) (model.Account, error)
}

type accountService struct {
	db *gorm.DB
	accountRepository repository.AccountRepository
	accountProfileRepository repository.AccountProfileRepository
}

func NewAccountService(
	db *gorm.DB,
	accountRepository repository.AccountRepository,
	accountProfileRepository repository.AccountProfileRepository,	
) AccountService {
	return &accountService{
		db: db,
		accountRepository: accountRepository,
		accountProfileRepository: accountProfileRepository,
	}
}

func (srv *accountService) Get(in input.Account) ([]model.Account, error) {
	return srv.accountRepository.Get(&model.Account{})
}

func (srv *accountService) GetOne(in input.Account) (model.Account, error) {
	return srv.accountRepository.GetOne(&model.Account{Id: in.Id})
}

func (srv *accountService) CreateOne(in input.Account) (model.Account, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.Account{}, err
	}

	var account model.Account
	err = srv.db.Transaction(func(tx *gorm.DB) error {
		account, err = srv.accountRepository.InsertTx(&model.Account{
			Name:     in.Name,
			Password: string(hashed),
			Role: helper.ACCOUNT_ROLE_NOMAL,
		}, tx)
		if err != nil {
			return err
		}
		_, err = srv.accountProfileRepository.InsertTx(&model.AccountProfile{
			AccountId: account.Id,
			DisplayName: in.Name,
		}, tx)
		if err != nil {
			return err
		}
		return nil
	})
	return account, err	
}

func (srv *accountService) UpdateOne(in input.Account) (model.Account, error) {
	account, err := srv.GetOne(in)
	if err != nil {
		return model.Account{}, err
	}
	account.Name = in.Name
	return srv.accountRepository.Update(&account)
}

func (srv *accountService) DeleteOne(in input.Account) error {
	return srv.accountRepository.Delete(&model.Account{Id: in.Id})
}

func (srv *accountService) Login(in input.Login) (model.Account, error) {
	account, err := srv.accountRepository.GetOne(&model.Account{Name: in.Name})
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			return model.Account{}, core.ErrUnauthorized
		}
		return model.Account{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(in.Password)); err != nil {
		return model.Account{}, core.ErrUnauthorized
	}
	return account, nil
}

func (srv *accountService) UpdatePassword(in input.UpdatePassword) (model.Account, error) {
	account, err := srv.accountRepository.GetOne(&model.Account{Id: in.Id})
	if err != nil {
		return model.Account{}, err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.Account{}, err
	}
	account.Password = string(hashed)
	return srv.accountRepository.Update(&account)
}