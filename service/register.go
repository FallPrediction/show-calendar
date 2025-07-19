package service

import (
	"errors"
	"souflair/models"
	"souflair/repository"
	"souflair/request"
	"souflair/utils"
)

type RegisterService struct {
	repository repository.UserRepository
}

func (s *RegisterService) Create(request *request.RegisterRequest) error {
	password, err := (&utils.Hash{}).HashEncoded(request.Password)
	if err != nil {
		return errors.New("hash 密碼失敗")
	}
	user := &models.User{
		Name:      request.Name,
		Password:  password,
		Email:     request.Email,
		Subscribe: true,
	}
	return s.repository.Create(user)
}

func NewRegisterService(repo repository.UserRepository) RegisterService {
	return RegisterService{repo}
}
