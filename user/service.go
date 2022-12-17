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

func (s *serviceImpl) GetById(ctx context.Context, id string) (*model.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.repository.GetById(ctx, &objectId)
}

func (s *serviceImpl) Update(ctx context.Context, id string, user *model.User) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	user.ID = objectId
	countUpdated, err := s.repository.Update(ctx, user)
	log.Printf("updated %d documents", countUpdated)
	return err
}

func (s *serviceImpl) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	countDeleted, err := s.repository.Delete(ctx, &objectId)
	log.Printf("deleted %d documents", countDeleted)
	return err
}
