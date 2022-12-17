package user

import (
	"fmt"
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/model"
)

type serviceImpl struct {
	users []model.User
}

func NewServiceImpl() interfaces.UserService {
	service := serviceImpl{}
	service.users = []model.User{
		{
			"123", "Eiji", 18, "eiji.ok@gmail.com", "zxjkdasfio", "rua x",
		},
		{
			"124", "Fabio", 28, "fabio.ok@gmail.com", "zxcv", "rua y",
		},
	}
	return &service
}

func (s *serviceImpl) List() ([]model.User, error) {
	return s.users, nil
}

func (s *serviceImpl) Save(user model.User) (model.User, error) {
	user.ID = "123"
	return user, nil
}

func (s *serviceImpl) GetById(id string) (model.User, error) {
	// read by user id
	fmt.Printf("getting User", id)
	return s.users[0], nil
}

func (s *serviceImpl) Update(user model.User) error {
	s.users[0] = user
	return nil
}

func (s *serviceImpl) Delete(id string) error {
	fmt.Printf("deleting User", id)
	s.users = s.users[1:]
	return nil
}
