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
	VerifyRegisterToken(b web.UserRegisterVerificationTokenRequest) (domain.User, error)
	VerifyForgotPasswordToken(b web.UserNewPasswordRequest) (domain.User, error)
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

func (s *authService) VerifyRegisterToken(request web.UserRegisterVerificationTokenRequest) (domain.User, error) {
	user := domain.User{}
	user, err := s.userRepository.FindById(uint(request.ID))
	if err != nil {
		return user, err
	}
	user.VerificationTime = request.VerificationTime
	return s.userRepository.Update(user), nil
}

func (s *authService) VerifyForgotPasswordToken(request web.UserNewPasswordRequest) (domain.User, error) {
	user := domain.User{}
	user, err := s.userRepository.FindById(uint(request.ID))
	if err != nil {
		return user, err
	}
	user.Password = request.Password
	return s.userRepository.Update(user), nil
}
