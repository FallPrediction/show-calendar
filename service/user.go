package service

import (
	"show-calendar/repository"
	"show-calendar/request"
)

type UserService struct {
	repository repository.UserRepository
}

func (service *UserService) LikeShow(request *request.UserLikeShowRequest, userId uint32) error {
	return service.repository.LikeShow(userId, request.ShowId)
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo}
}
