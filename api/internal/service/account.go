package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"lv99/internal/core"
	"lv99/internal/dto/input"
	"lv99/internal/helper"
	"lv99/internal/model"
	"lv99/internal/repository"
)

type AccountService interface {
	GetOne(in input.Account) (model.Account, error)
	DeleteOne(in input.Account) error
	UpdateOne(in input.Account) (model.Account, error)
	Login(in input.Login) (model.Account, error)
	Signup(in input.Signup) (model.Account, error)
}

type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) AccountService {
	return &accountService{
		accountRepository: accountRepository,
	}
}

func (srv *accountService) GetOne(in input.Account) (model.Account, error) {
	account, err := srv.accountRepository.GetOne(&model.Account{AccountId: in.AccountId})
	return account, err
}

func (srv *accountService) UpdateOne(in input.Account) (model.Account, error) {
	account, err := srv.GetOne(in)
	if err != nil {
		return model.Account{}, err
	}

	if in.AccountName != "" {
		account.AccountName = in.AccountName
	}
	if in.AccountPassword != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(in.AccountPassword), bcrypt.DefaultCost)
		if err != nil {
			return model.Account{}, err
		}
		account.AccountPassword = string(hashed)
	}
	account, err = srv.accountRepository.Update(&account)
	return account, err
}

func (srv *accountService) DeleteOne(in input.Account) error {
	err := srv.accountRepository.Delete(&model.Account{AccountId: in.AccountId})
	return err
}

func (srv *accountService) Login(in input.Login) (model.Account, error) {
	account, err := srv.accountRepository.GetOne(&model.Account{AccountName: in.AccountName})
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			return model.Account{}, core.ErrUnauthorized
		}
		return model.Account{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(account.AccountPassword), []byte(in.AccountPassword)); err != nil {
		return model.Account{}, core.ErrUnauthorized
	}
	return account, nil
}

func (srv *accountService) Signup(in input.Signup) (model.Account, error) {
	if _, err := srv.accountRepository.GetOne(&model.Account{AccountName: in.AccountName}); err == nil {
		return model.Account{}, core.ErrConflict
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(in.AccountPassword), bcrypt.DefaultCost)
	if err != nil {
		return model.Account{}, err
	}

	account := model.Account{
		AccountName:     in.AccountName,
		AccountPassword: string(hashed),
		AccountRole: helper.ACCOUNT_ROLE_NOMAL,
	}

	account, err = srv.accountRepository.Insert(&account)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}
