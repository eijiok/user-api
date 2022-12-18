package user

import (
	"context"
	"github.com/eijiok/user-api/dto"
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

func (s *serviceImpl) List(ctx context.Context) ([]dto.User, error) {

	list, err := s.repository.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]dto.User, len(list))
	for i, user := range list {
		dtoUser := dto.User{}
		dtoUser.FromUserModel(&user)
		result[i] = dtoUser
	}

	return result, err
}

func (s *serviceImpl) Save(ctx context.Context, user *model.User) (*dto.User, error) {
	id, err := s.repository.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	dtoUser := dto.User{}
	dtoUser.FromUserModel(user)
	dtoUser.ID = *id

	return &dtoUser, nil
}

func (s *serviceImpl) GetById(ctx context.Context, objectID *primitive.ObjectID) (*dto.User, error) {
	user, err := s.repository.GetById(ctx, objectID)
	dtoUser := dto.User{}
	dtoUser.FromUserModel(user)
	return &dtoUser, err
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
