package user

import (
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/model"
)

type serviceImpl struct {
	repository interfaces.UserRepository
}

func NewServiceImpl(repository interfaces.UserRepository) interfaces.UserService {
	return &serviceImpl{
		repository: repository,
	}
}

func (s *serviceImpl) List() ([]model.User, error) {
	return s.repository.List()
}

func (s *serviceImpl) Save(user model.User) (model.User, error) {
	return s.repository.Save(user)
}

func (s *serviceImpl) GetById(id string) (model.User, error) {
	return s.repository.GetById(id)
}

func (s *serviceImpl) Update(user model.User) error {
	return s.repository.Update(user)
}

func (s *serviceImpl) Delete(id string) error {
	return s.repository.Delete(id)
}
