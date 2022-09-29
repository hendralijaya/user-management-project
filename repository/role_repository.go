package repository

import (
	"errors"
	"hendralijaya/user-management-project/model/domain"

	"gorm.io/gorm"
)

type RoleRepository interface {
	All() []domain.Role
	Create(r domain.Role) domain.Role
	Update(r domain.Role) domain.Role
	Delete(r domain.Role)
	FindById(id uint64) (domain.Role, error)
}

type RoleConnection struct {
	connection *gorm.DB
}

func NewRoleRepository(connection *gorm.DB) RoleRepository {
	return &RoleConnection {connection: connection}
}

func (c *RoleConnection) All() []domain.Role {
	var roles []domain.Role
	c.connection.Find(&roles)
	return roles
}

func (c *RoleConnection) Create(r domain.Role) domain.Role {
	c.connection.Save(&r)
	return r
}

func (c *RoleConnection) Update(r domain.Role) domain.Role {
	c.connection.Save(&r)
	return r
}

func (c *RoleConnection) Delete(r domain.Role) {
	c.connection.Delete(&r)
}

func (c *RoleConnection) FindById(id uint64) (domain.Role, error) {
	var role domain.Role
	c.connection.Find(&role, "id = ?", id)
	if role.ID == 0 {
		return role, errors.New("role id not found")
	}
	return role, nil
}