package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"lv99/internal/core"
	"lv99/internal/helper"
	accountModule "lv99/internal/module/account"
	profileModule "lv99/internal/module/account_profile"
)

type Usecase interface {
	Signup(in SignupDto, db *gorm.DB) (accountModule.Account, error)
	Login(in LoginDto, db *gorm.DB) (accountModule.Account, string, string, error)
	Refresh(refreshToken string, db *gorm.DB) (core.AuthPayload, string, error)
	UpdatePassword(in UpdatePasswordDto, db *gorm.DB) error
}

type usecase struct {
	accountService        accountModule.Service
	accountProfileService profileModule.Service
}

func NewUsecase(
	accountService accountModule.Service,
	accountProfileService profileModule.Service,
) Usecase {
	return &usecase{
		accountService:        accountService,
		accountProfileService: accountProfileService,
	}
}

func (srv *usecase) Signup(in SignupDto, db *gorm.DB) (accountModule.Account, error) {
	var account accountModule.Account
	err := db.Transaction(func(tx *gorm.DB) error {
		hashed, err := hashPassword(in.Password)
		if err != nil {
			return err
		}

		account, err = srv.accountService.CreateOne(accountModule.Account{
			Name:     in.Name,
			Password: string(hashed),
			Role:     helper.ACCOUNT_ROLE_NOMAL,
		}, tx)

		if err != nil {
			return err
		}

		_, err = srv.accountProfileService.CreateOne(profileModule.AccountProfile{
			AccountId:   account.Id,
			DisplayName: account.Name,
			Bio:         "",
			AvatarURL:   "",
		}, tx)
		return err
	})

	return account, err
}

func (srv *usecase) Login(in LoginDto, db *gorm.DB) (accountModule.Account, string, string, error) {
	acct, err := srv.accountService.GetOne(accountModule.Account{Name: in.Name}, db)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			return accountModule.Account{}, "", "", core.ErrUnauthorized
		}
		return accountModule.Account{}, "", "", err
	}

	if err = verifyPassword(acct.Password, in.Password); err != nil {
		return accountModule.Account{}, "", "", core.ErrUnauthorized
	}

	accessToken, err := core.Auth.CreateAccessToken(core.AuthPayload{
		AccountId:   acct.Id,
		AccountName: acct.Name,
		AccountRole: acct.Role,
	})
	if err != nil {
		return accountModule.Account{}, "", "", err
	}

	refreshToken, err := core.Auth.CreateRefreshToken(core.AuthPayload{
		AccountId:   acct.Id,
		AccountName: acct.Name,
		AccountRole: acct.Role,
	})
	if err != nil {
		return accountModule.Account{}, "", "", err
	}
	return acct, accessToken, refreshToken, nil
}

func (srv *usecase) Refresh(refreshToken string, db *gorm.DB) (core.AuthPayload, string, error) {
	payload, err := core.Auth.VerifyRefreshToken(refreshToken)
	if err != nil {
		return core.AuthPayload{}, "", core.NewAppError("invalid or expired refresh token", core.ErrCodeUnauthorized)
	}

	accessToken, err := core.Auth.CreateAccessToken(core.AuthPayload{
		AccountId:   payload.AccountId,
		AccountName: payload.AccountName,
		AccountRole: payload.AccountRole,
	})

	return payload, accessToken, err
}

func (srv *usecase) UpdatePassword(in UpdatePasswordDto, db *gorm.DB) error {
	acct, err := srv.accountService.GetOne(accountModule.Account{Id: in.Id}, db)
	if err != nil {
		return err
	}
	if err = verifyPassword(acct.Password, in.OldPassword); err != nil {
		return core.ErrBadRequest
	}

	hashed, err := hashPassword(in.NewPassword)
	if err != nil {
		return err
	}
	acct.Password = string(hashed)
	_, err = srv.accountService.UpdateOne(acct, db)
	return err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func verifyPassword(hashedPassword, plainPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return err
	}
	return nil
}
