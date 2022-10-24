//go:build wireinject
// +build wireinject

package injector

import (
	"hendralijaya/user-management-project/controller"
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/middleware"
	"hendralijaya/user-management-project/repository"
	"hendralijaya/user-management-project/service"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var jwtMiddlewareSet = wire.NewSet(
	service.NewJWTService,
	middleware.NewAuthorizeJWTMiddleware,
)

var adminMiddlewareSet = wire.NewSet(
	logrus.New,
	helper.NewLog,
	repository.NewUserRepository,
	service.NewJWTService,
	service.NewUserService,
	middleware.NewIsAdminMiddleware,
)

var roleSet = wire.NewSet(
	logrus.New,
	helper.NewLog,
	repository.NewRoleRepository,
	service.NewRoleService,
	service.NewJWTService,
	controller.NewRoleController,
)

var userSet = wire.NewSet(
	logrus.New,
	helper.NewLog,
	repository.NewUserRepository,
	service.NewUserService,
	service.NewJWTService,
	controller.NewUserController,
)

var authSet = wire.NewSet(
	logrus.New,
	helper.NewLog,
	repository.NewUserRepository,
	service.NewAuthService,
	service.NewUserService,
	service.NewJWTService,
	controller.NewAuthController,
)

func InitRole(db *gorm.DB) controller.RoleController {
	wire.Build(
		roleSet,
	)
	return nil
}

func InitUser(db *gorm.DB, mongoDB *mongo.Client) controller.UserController {
	wire.Build(
		userSet,
	)
	return nil
}

func InitAuth(db *gorm.DB, mongoDB *mongo.Client) controller.AuthController {
	wire.Build(
		authSet,
	)
	return nil
}

func InitJWTMiddleware() middleware.AuthorizeJWTMiddleware {
	wire.Build(
		jwtMiddlewareSet,
	)
	return nil
}

func InitAdminMiddleware(db *gorm.DB, mongoDB *mongo.Client) middleware.IsAdminMiddleware {
	wire.Build(
		adminMiddlewareSet,
	)
	return nil
}
