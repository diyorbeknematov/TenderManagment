package service

import (
	"fmt"
	"tender/model"
)

func (s *Service) Registration(user model.UserRegisterReq) (*model.UserRegisterResp, error) {
	resp, err := s.Storage.RegistrationRepository().CreateUser(user)

	if err != nil {
		s.Log.Error(fmt.Sprintf("Register crudda xatolik bor: %v", err))
		return nil, err
	}

	return resp, nil
}

func (s *Service) GetUserByEmail(email string) (*model.GetUser, error) {
	resp, err := s.Storage.RegistrationRepository().GetUserByEmail(email)

	if err != nil {
		s.Log.Error(fmt.Sprintf("Userni email orqali olish crudida xatolik bor: %v", err))
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
