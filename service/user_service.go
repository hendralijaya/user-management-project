package service

import (
	"hendralijaya/user-management-project/model/domain"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	All() []domain.User
	Create(b web.UserCreateRequest) (domain.User, error)
	FindById(id uint) (domain.User, error)
	Update(b web.UserUpdateRequest) (domain.User, error)
	FindByEmail(email string) domain.User
	FindByUsername(username string) domain.User
	Delete(id uint) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (service *userService) All() []domain.User {
	return service.userRepository.All()
}

func (service *userService) Create(request web.UserCreateRequest) (domain.User, error) {
	user := domain.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&request))

	if err != nil {
		return user, err
	}

	_, err = service.userRepository.IsDuplicateEmail(request.Email)
	if err != nil {
		return user, err
	}
	return service.userRepository.Create(user), nil
}

func (service *userService) Update(b web.UserUpdateRequest) (domain.User, error) {
	user := domain.User{}
	_, err := service.userRepository.FindById(uint(b.ID))
	if err != nil {
		return user, err
	}
	err = smapping.FillStruct(&user, smapping.MapFields(&b))
	if err != nil {
		return user, err
	}
	if b.Password != "" {
		return service.userRepository.Update(user, true), nil
	} else {
		return service.userRepository.Update(user, false), nil
	}
}

func (service *userService) FindById(id uint) (domain.User, error) {
	user, err := service.userRepository.FindById(uint(id))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *userService) FindByEmail(email string) domain.User {
	user := service.userRepository.FindByEmail(email)
	return user
}

func (service *userService) FindByUsername(username string) domain.User {
	user := service.userRepository.FindByUsername(username)
	return user
}

func (service *userService) Delete(id uint) error {
	user, err := service.userRepository.FindById(uint(id))
	if err != nil {
		return err
	}
	service.userRepository.Delete(user)
	return nil
}
