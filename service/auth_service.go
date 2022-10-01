package service

import (
	"hendralijaya/user-management-project/model/domain"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/repository"
)

type AuthService interface {
	Login(b web.UserLoginRequest) (domain.User, error)
	Register(b web.UserRegisterRequest) (domain.User, error)
	VerifyRegisterToken(b web.UserRegisterVerificationRequest) (domain.User, error)
	VerifyForgotPasswordToken(b web.UserUpdateRequest) (domain.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *authService) Login(u web.UserLoginRequest) (domain.User, error) {
	user, err := s.userRepository.VerifyCredential(u.Username, u.Email, u.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
