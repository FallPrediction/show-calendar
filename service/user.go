package service

import (
	"show-calendar/repository"
	"show-calendar/request"
	"show-calendar/utils"
)

type UserService struct {
	repository repository.UserRepository
}

func (service *UserService) LikeShow(request *request.UserLikeShowRequest, userId uint32) error {
	return service.repository.LikeShow(userId, request.ShowId)
}

func (service *UserService) Unsubscribe(token string) error {
	aes := utils.NewAes()
	return service.repository.Unsubscribe(aes.Decrypt(token))
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo}
}
