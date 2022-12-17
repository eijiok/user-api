package user

import (
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type repositoryImpl struct {
	collection *mongo.Collection
}

func (r repositoryImpl) List() ([]model.User, error) {
	//r.collection.Find()
	panic("implement me")
}

func (r repositoryImpl) GetById(id string) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repositoryImpl) Save(user model.User) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repositoryImpl) Update(user model.User) error {
	//TODO implement me
	panic("implement me")
}

func (r repositoryImpl) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(database *mongo.Database) interfaces.UserRepository {
	return &repositoryImpl{
		collection: database.Collection("user"),
	}
}
