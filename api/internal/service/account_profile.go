package service

import (
	"lv99/internal/dto/input"
	"lv99/internal/model"
	"lv99/internal/repository"
)

type AccountProfileService interface {
	Get(in input.AccountProfile) ([]model.AccountProfile, error)
	GetOne(in input.AccountProfile) (model.AccountProfile, error)
	CreateOne(in input.AccountProfile) (model.AccountProfile, error)
	UpdateOne(in input.AccountProfile) (model.AccountProfile, error)
	DeleteOne(in input.AccountProfile) error
}

type accountProfileService struct {
	accountProfileRepository repository.AccountProfileRepository
}

func NewAccountProfileService(accountProfileRepository repository.AccountProfileRepository) AccountProfileService {
	return &accountProfileService{
		accountProfileRepository: accountProfileRepository,
	}
}

func (srv *accountProfileService) Get(in input.AccountProfile) ([]model.AccountProfile, error) {
	return srv.accountProfileRepository.Get(&model.AccountProfile{})
}

func (srv *accountProfileService) GetOne(in input.AccountProfile) (model.AccountProfile, error) {
	return srv.accountProfileRepository.GetOne(&model.AccountProfile{AccountId: in.AccountId})
}

func (srv *accountProfileService) CreateOne(in input.AccountProfile) (model.AccountProfile, error) {
	return srv.accountProfileRepository.Insert(&model.AccountProfile{
		AccountId:   in.AccountId,
		DisplayName: in.DisplayName,
		Bio:         in.Bio,
		AvatarURL:   in.AvatarURL,
	})
}

func (srv *accountProfileService) UpdateOne(in input.AccountProfile) (model.AccountProfile, error) {
	profile, err := srv.GetOne(in)
	if err != nil {
		return model.AccountProfile{}, err
	}

	profile.DisplayName = in.DisplayName
	profile.Bio = in.Bio
	profile.AvatarURL = in.AvatarURL

	return srv.accountProfileRepository.Update(&profile)
}

func (srv *accountProfileService) DeleteOne(in input.AccountProfile) error {
	return srv.accountProfileRepository.Delete(&model.AccountProfile{AccountId: in.AccountId})
}
