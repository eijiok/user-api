package user

import (
	"github.com/eijiok/user-api/db"
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/middleware"
	"github.com/eijiok/user-api/security"
	"go.mongodb.org/mongo-driver/mongo"
)

var factoryInstance interfaces.UserFactory = nil

type factoryImpl struct {
	database   *mongo.Database
	service    interfaces.UserService
	controller interfaces.UserController
	repository interfaces.UserRepository
	router     interfaces.UserRouter
}

func (factory *factoryImpl) GetController() interfaces.UserController {
	if factory.controller == nil {
		factory.controller = NewControllerImpl(factory.GetService())
	}
	return factory.controller
}

func (factory *factoryImpl) GetService() interfaces.UserService {
	if factory.service == nil {
		factory.service = NewServiceImpl(factory.GetRepository(), security.HashPassword)
	}
	return factory.service
}

func (factory *factoryImpl) GetRepository() interfaces.UserRepository {
	if factory.repository == nil {
		factory.repository = NewUserRepository(factory.database)
	}
	return factory.repository
}

func (factory *factoryImpl) GetRouter() interfaces.UserRouter {
	if factory.router == nil {
		factory.router = NewUserRouter(
			factory.GetController(),
			[]interfaces.MiddlewareFunc{middleware.CorsMiddleware, middleware.AuthMiddleware},
			[]interfaces.MiddlewareFunc{middleware.CorsMiddleware},
		)
	}
	return factory.router
}

func GetFactory(mongoConfig *db.MongoConfig) interfaces.UserFactory {
	if factoryInstance == nil {
		factoryInstance = &factoryImpl{
			database: mongoConfig.Database,
		}
	}
	return factoryInstance
}
