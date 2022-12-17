package user

import (
	"context"
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repositoryImpl struct {
	collection *mongo.Collection
}

func (r repositoryImpl) List(ctx context.Context) ([]model.User, error) {
	filter := bson.D{{}}
	cursor, err := r.collection.Find(ctx, filter)
	var results []model.User
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, err
}

func (r repositoryImpl) GetById(ctx context.Context, id *primitive.ObjectID) (*model.User, error) {
	filter := bson.D{{"_id", id}}

	var result model.User
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r repositoryImpl) Save(ctx context.Context, user *model.User) (*primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return &id, nil
}

func (r repositoryImpl) Update(ctx context.Context, user *model.User) (int64, error) {
	filter := bson.D{{"_id", user.ID}}
	update := bson.D{{"$set", *user}}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	return result.ModifiedCount, err
}

func (r repositoryImpl) Delete(ctx context.Context, id *primitive.ObjectID) (int64, error) {

	filter := bson.D{{"_id", id}}
	result, err := r.collection.DeleteOne(ctx, filter)
	return result.DeletedCount, err
}

func NewUserRepository(database *mongo.Database) interfaces.UserRepository {
	return &repositoryImpl{
		collection: database.Collection("user"),
	}
}
