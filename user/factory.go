package user

import (
	"github.com/eijiok/user-api/db"
	"github.com/eijiok/user-api/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

var factoryInstance interfaces.UserFactory = nil

type factoryImpl struct {
	database   *mongo.Database
	service    interfaces.UserService
	controller interfaces.UserController
	repository interfaces.UserRepository
}

func (factory *factoryImpl) GetController() interfaces.UserController {
	if factory.controller == nil {
		factory.controller = NewControllerImpl(factory.GetService())
	}
	return factory.controller
}

func (factory *factoryImpl) GetService() interfaces.UserService {
	if factory.service == nil {
		factory.service = NewServiceImpl(factory.GetRepository())
	}
	return factory.service
}

func (factory *factoryImpl) GetRepository() interfaces.UserRepository {
	if factory.repository == nil {
		factory.repository = NewUserRepository(factory.database)
	}
	return factory.repository
}

func GetFactory(mongoConfig *db.MongoConfig) interfaces.UserFactory {
	if factoryInstance == nil {
		factoryInstance = &factoryImpl{
			database: mongoConfig.Database,
		}
	}
	return factoryInstance
}