package service

import (
	"hendralijaya/user-management-project/model/domain"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	All() []domain.User
	VerifyCredential(b web.UserLoginRequest) (interface{}, error)
	Create(b web.UserRegisterRequest) (domain.User, error)
	FindById(id uint64) (domain.User, error)
	Update(b domain.User) (domain.User, error)
	FindByEmail(email string) domain.User
	// Logout(u web.UserLogoutRequest) (domain.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) All() []domain.User {
	return s.userRepository.All()
}

func (s *userService) VerifyCredential(u web.UserLoginRequest) (interface{}, error) {
	user, err := s.userRepository.VerifyCredential(u.Email, u.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) Create(request web.UserRegisterRequest) (domain.User, error) {
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

func (s *userService) Update(b domain.User) (domain.User, error) {
	_, err := s.userRepository.FindById(b.Id)
	if err != nil {
		return b, err
	}
	return s.userRepository.Update(b), nil
}

func (s *userService) FindById(id uint64) (domain.User, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) FindByEmail(email string) domain.User {
	user := s.userRepository.FindByEmail(email)
	return user
}
