package user

import (
	"context"
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type serviceImpl struct {
	repository interfaces.UserRepository
}

func NewServiceImpl(repository interfaces.UserRepository) interfaces.UserService {
	return &serviceImpl{
		repository: repository,
	}
}

func (s *serviceImpl) List(ctx context.Context) ([]model.User, error) {
	return s.repository.List(ctx)
}

func (s *serviceImpl) Save(ctx context.Context, user *model.User) (*model.User, error) {
	id, err := s.repository.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = *id
	user.Password = ""
	return user, nil
}

func (s *serviceImpl) GetById(ctx context.Context, objectID *primitive.ObjectID) (*model.User, error) {
	return s.repository.GetById(ctx, objectID)
}

func (s *serviceImpl) Update(ctx context.Context, objectID *primitive.ObjectID, user *model.User) error {
	user.ID = *objectID
	countUpdated, err := s.repository.Update(ctx, user)
	log.Printf("updated %d documents", countUpdated)
	return err
}

func (s *serviceImpl) Delete(ctx context.Context, objectId *primitive.ObjectID) error {
	countDeleted, err := s.repository.Delete(ctx, objectId)
	log.Printf("deleted %d documents", countDeleted)
	return err
}
