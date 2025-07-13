package account_profile

import (
	profileModule "lv99/internal/module/account_profile"

	"gorm.io/gorm"
)

type Usecase interface {
	Get(in GetDto, db *gorm.DB) ([]profileModule.AccountProfile, error)
	GetOne(in GetOneDto, db *gorm.DB) (profileModule.AccountProfile, error)
	CreateOne(in CreateOneDto, db *gorm.DB) (profileModule.AccountProfile, error)
	UpdateOne(in UpdateOneDto, db *gorm.DB) (profileModule.AccountProfile, error)
	DeleteOne(in DeleteOneDto, db *gorm.DB) error
}

type usecase struct {
	accountProfileService profileModule.Service
}

func NewUsecase(accountProfileService profileModule.Service) Usecase {
	return &usecase{
		accountProfileService: accountProfileService,
	}
}

func (uc *usecase) Get(in GetDto, db *gorm.DB) ([]profileModule.AccountProfile, error) {
	return uc.accountProfileService.Get(profileModule.AccountProfile{}, db)
}

func (uc *usecase) GetOne(in GetOneDto, db *gorm.DB) (profileModule.AccountProfile, error) {
	return uc.accountProfileService.GetOne(profileModule.AccountProfile{AccountId: in.AccountId}, db)
}

func (uc *usecase) CreateOne(in CreateOneDto, db *gorm.DB) (profileModule.AccountProfile, error) {
	return uc.accountProfileService.CreateOne(profileModule.AccountProfile{
		AccountId:   in.AccountId,
		DisplayName: in.DisplayName,
		Bio:         in.Bio,
		AvatarURL:   in.AvatarURL,
	}, db)
}

func (uc *usecase) UpdateOne(in UpdateOneDto, db *gorm.DB) (profileModule.AccountProfile, error) {
	return uc.accountProfileService.UpdateOne(profileModule.AccountProfile{
		AccountId:   in.AccountId,
		DisplayName: in.DisplayName,
		Bio:         in.Bio,
		AvatarURL:   in.AvatarURL,
	}, db)
}

func (uc *usecase) DeleteOne(in DeleteOneDto, db *gorm.DB) error {
	return uc.accountProfileService.DeleteOne(profileModule.AccountProfile{AccountId: in.AccountId}, db)
}
