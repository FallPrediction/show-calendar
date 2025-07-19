package service

import (
	"souflair/repository"
	"souflair/request"
	"souflair/utils"
)

type UserService struct {
	repository repository.UserRepository
}

func (s *UserService) LikeShow(request *request.UserLikeShowRequest, userId uint32) error {
	return s.repository.LikeShow(userId, request.ShowId)
}

func (s *UserService) Unsubscribe(token string) error {
	aes := utils.NewAes()
	return s.repository.Unsubscribe(aes.Decrypt(token))
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo}
}
