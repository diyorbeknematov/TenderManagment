package service

import (
	"fmt"
	"tender/model"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Registration(user model.UserRegisterReq) (*model.UserRegisterResp, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Xatolik passwordni hashlashda: %v", err))
		return nil, err
	}

	user.Password = string(hashPass)

	resp, err := s.Storage.RegistrationRepository().CreateUser(user)

	if err != nil {
		s.Log.Error(fmt.Sprintf("Register crudda xatolik bor: %v", err))
		return nil, err
	}

	return resp, nil
}

func (s *Service) GetUserByUsername(login model.LoginUser) (*model.GetUser, error) {
	resp, err := s.Storage.RegistrationRepository().GetUserByUsername(login.Username)

	if err != nil {
		s.Log.Error(fmt.Sprintf("Userni email orqali olish crudida xatolik bor: %v", err))
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(login.Password))
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error passwordni tekshirishda xatolik bor: %v", err))
		return nil, err
	}

	return resp, nil
}

func (s *Service) IsUserExists(exists model.IsUserExists) (bool, error) {
	isTrue, err := s.Storage.RegistrationRepository().IsUserExists(exists.Email, exists.Username)

	if err != nil {
		s.Log.Error(fmt.Sprintf("Userni bor yoki yo'qligini tekshirishda xatolik bor: %v", err))
		return false, err
	}

	return isTrue, err
}
