package service

import (
	"hendralijaya/user-management-project/model/domain"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/repository"

	"github.com/mashingan/smapping"
)

type RoleService interface {
	All() []domain.Role
	Create(r web.RoleCreateRequest) (domain.Role, error)
	FindById(id uint) (domain.Role, error)
	Update(r web.RoleUpdateRequest) (domain.Role, error)
	Delete(id uint) error
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) RoleService {
	return &roleService{roleRepository: roleRepository}
}

func (service *roleService) All() []domain.Role {
	return service.roleRepository.All()
}

func (service *roleService) Create(r web.RoleCreateRequest) (domain.Role, error) {
	role := domain.Role{}
	err := smapping.FillStruct(&role, smapping.MapFields(&r))
	if err != nil {
		return role, err
	}
	return service.roleRepository.Create(role), nil
}

func (service *roleService) FindById(id uint) (domain.Role, error) {
	role, err := service.roleRepository.FindById(uint64(id))
	if err != nil {
		return role, err
	}
	return role, nil
}

func (service *roleService) Update(r web.RoleUpdateRequest) (domain.Role, error) {
	role := domain.Role{}
	_, err := service.roleRepository.FindById(r.ID)
	if err != nil {
		return role, err
	}
	err = smapping.FillStruct(&role, smapping.MapFields(&r))
	if err != nil {
		return role, err
	}
	return service.roleRepository.Update(role), nil
}

func (service *roleService) Delete(id uint) error {
	role, err := service.roleRepository.FindById(uint64(id))
	if err != nil {
		return err
	}
	service.roleRepository.Delete(role)
	return nil
}
