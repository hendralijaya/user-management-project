package service

import (
	"hendralijaya/user-management-project/model/domain"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/repository"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	Login(b web.UserLoginRequest) (domain.User, error)
	Register(b web.UserRegisterRequest) (domain.User, error)
	ForgotPassword(b web.UserForgotPasswordRequest) (domain.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (s *authService) Login(u web.UserLoginRequest) (domain.User, error) {
	user, err := s.userRepository.VerifyCredential(u.Username, u.Email, u.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *authService) Register(request web.UserRegisterRequest) (domain.User, error) {
	user := domain.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&request))

	if err != nil {
		return user, err
	}

	_, err = s.userRepository.IsDuplicateEmail(request.Email)
	if err != nil {
		return user, err
	}
	return s.userRepository.Create(user), nil
}

func (s *authService) ForgotPassword(request web.UserForgotPasswordRequest) (domain.User, error) {
	user := domain.User{}
	 return user, nil
}
