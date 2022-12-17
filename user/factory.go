package user

import (
	"github.com/eijiok/user-api/interfaces"
)

var factoryInstance interfaces.Factory = nil

type factoryImpl struct {
	service    interfaces.UserService
	controller interfaces.UserController
}

func (factory *factoryImpl) GetController() interfaces.UserController {
	if factory.controller == nil {
		factory.controller = NewControllerImpl(factory.GetService())
	}
	return factory.controller
}

func (factory *factoryImpl) GetService() interfaces.UserService {
	if factory.service == nil {
		factory.service = NewServiceImpl()
	}
	return factory.service
}

func GetFactory() interfaces.Factory {
	if factoryInstance == nil {
		factoryInstance = &factoryImpl{}
	}
	return factoryInstance
}
