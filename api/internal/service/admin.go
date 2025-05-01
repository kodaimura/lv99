package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"lv99/internal/core"
	"lv99/internal/dto/input"
	"lv99/internal/model"
	"lv99/internal/repository"
)

type AdminService interface {
	GetOne(in input.Admin) (model.Admin, error)
	DeleteOne(in input.Admin) error
	UpdateOne(in input.Admin) (model.Admin, error)
	Login(in input.AdminLogin) (model.Admin, error)
	Signup(in input.AdminSignup) (model.Admin, error)
}

type adminService struct {
	adminRepository repository.AdminRepository
}

func NewAdminService(adminRepository repository.AdminRepository) AdminService {
	return &adminService{
		adminRepository: adminRepository,
	}
}

func (srv *adminService) GetOne(in input.Admin) (model.Admin, error) {
	admin, err := srv.adminRepository.GetOne(&model.Admin{AdminId: in.AdminId})
	return admin, err
}

func (srv *adminService) UpdateOne(in input.Admin) (model.Admin, error) {
	admin, err := srv.GetOne(in)
	if err != nil {
		return model.Admin{}, err
	}

	if in.AdminName != "" {
		admin.AdminName = in.AdminName
	}
	if in.AdminPassword != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(in.AdminPassword), bcrypt.DefaultCost)
		if err != nil {
			return model.Admin{}, err
		}
		admin.AdminPassword = string(hashed)
	}
	admin, err = srv.adminRepository.Update(&admin)
	return admin, err
}

func (srv *adminService) DeleteOne(in input.Admin) error {
	err := srv.adminRepository.Delete(&model.Admin{AdminId: in.AdminId})
	return err
}

func (srv *adminService) Login(in input.AdminLogin) (model.Admin, error) {
	admin, err := srv.adminRepository.GetOne(&model.Admin{AdminName: in.AdminName})
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			return model.Admin{}, core.ErrUnauthorized
		}
		return model.Admin{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(admin.AdminPassword), []byte(in.AdminPassword)); err != nil {
		return model.Admin{}, core.ErrUnauthorized
	}
	return admin, nil
}

func (srv *adminService) Signup(in input.AdminSignup) (model.Admin, error) {
	if _, err := srv.adminRepository.GetOne(&model.Admin{AdminName: in.AdminName}); err == nil {
		return model.Admin{}, core.ErrConflict
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(in.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return model.Admin{}, err
	}

	admin := model.Admin{
		AdminName:     in.AdminName,
		AdminPassword: string(hashed),
	}

	admin, err = srv.adminRepository.Insert(&admin)
	if err != nil {
		return model.Admin{}, err
	}

	return admin, nil
}
