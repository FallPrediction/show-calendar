package service

import (
	custom_errors "souflair/errors"
	"souflair/initialize"
	"souflair/repository"
	"souflair/request"
	"souflair/utils"
	"time"

	"github.com/matthewhartstonge/argon2"
)

type AuthenticateService struct {
	repostiroy repository.UserRepository
}

func (s *AuthenticateService) Login(request *request.LoginRequest) (string, time.Time, error) {
	logger := initialize.NewLogger()
	user, err := s.repostiroy.GetByEmail(request.Email)
	if err != nil {
		return "", time.Time{}, err
	}
	ok, err := argon2.VerifyEncoded([]byte(request.Password), []byte(user.Password))
	if err != nil {
		logger.Error("Encode password failed. ", err)
		return "", time.Time{}, custom_errors.ErrPasswordIncorrect
	} else if !ok {
		return "", time.Time{}, custom_errors.ErrPasswordIncorrect
	}
	return (&utils.Jwt{}).CreateUserToken(&user)
}

func NewAuthenticateService(repostiroy repository.UserRepository) AuthenticateService {
	return AuthenticateService{repostiroy: repostiroy}
}
