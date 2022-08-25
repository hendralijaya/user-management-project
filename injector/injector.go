//go:build wireinject
// +build wireinject

package injector

import (
	"hendralijaya/user-management-project/controller"
	"hendralijaya/user-management-project/middleware"
	"hendralijaya/user-management-project/repository"
	"hendralijaya/user-management-project/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var jwtMiddlewareSet = wire.NewSet(
	service.NewJWTService,
	middleware.NewAuthorizeJWTMiddleware,
)

var adminMiddlewareSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewJWTService,
	service.NewUserService,
	middleware.NewIsAdminMiddleware,
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	service.NewJWTService,
	controller.NewUserController,
)

var authSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	service.NewJWTService,
	controller.NewAuthController,
)

func InitUser(db *gorm.DB) (controller.UserController){
	wire.Build(
		userSet,
	)
	return nil
}

func InitAuth(db *gorm.DB) controller.AuthController{
	wire.Build(
		authSet,
	)
	return nil
}

func InitJWTMiddleware() middleware.AuthorizeJWTMiddleware{
	wire.Build(
		jwtMiddlewareSet,
	)
	return nil
}

func InitAdminMiddleware(db *gorm.DB) middleware.IsAdminMiddleware{
	wire.Build(
		adminMiddlewareSet,
	)
	return nil
}