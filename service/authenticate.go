package service

import (
	custom_errors "show-calendar/errors"
	"show-calendar/initialize"
	"show-calendar/repository"
	"show-calendar/request"
	"show-calendar/utils"

	"github.com/matthewhartstonge/argon2"
)

type AuthenticateService struct {
	repostiroy repository.UserRepository
}

func (service *AuthenticateService) Login(request *request.LoginRequest) (string, error) {
	logger := initialize.NewLogger()
	user, err := service.repostiroy.GetByEmail(request.Email)
	if err != nil {
		return "", err
	}
	ok, err := argon2.VerifyEncoded([]byte(request.Password), []byte(user.Password))
	if err != nil {
		logger.Error("Encode password failed. ", err)
		return "", custom_errors.ErrPasswordIncorrect
	} else if !ok {
		return "", custom_errors.ErrPasswordIncorrect
	}
	return (&utils.Jwt{}).CreateUserToken(&user)
}

func NewAuthenticateService(repostiroy repository.UserRepository) AuthenticateService {
	return AuthenticateService{repostiroy: repostiroy}
}
