package user

import (
	"context"
	"errors"
	"github.com/eijiok/user-api/dto"
	mocks "github.com/eijiok/user-api/mocks/interfaces"
	"github.com/eijiok/user-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
	"time"
)

func Test_serviceImpl_GetById(t *testing.T) {
	type fields struct {
		repository   *mocks.UserRepository
		hashPassword func(password string) (string, error)
		timeService  *mocks.TimeService
	}
	type args struct {
		ctx      context.Context
		objectID *primitive.ObjectID
	}
	mockContext := context.Background()
	id := primitive.NewObjectID()
	name := "Eiji Okuda"
	address := "Rua das Flores"
	email := "eiji.ok@gmail.com"
	createdAt := time.Now()
	updatedAt := createdAt.Add(time.Second * 3)

	err := errors.New("Error x")

	birthday := time.Date(1981, 4, 14, 10, 34, 58, 0, time.UTC)
	user := model.User{
		ID:        id,
		Name:      name,
		Birthday:  birthday,
		Email:     email,
		Password:  "asdfpoiu",
		Address:   address,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	expectedUser := dto.UserResponse{
		ID:        id,
		Name:      name,
		Address:   address,
		Birthday:  birthday,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.UserResponse
		wantErr bool
	}{
		{
			name: "test GetById error",
			fields: fields{
				repository: (func() *mocks.UserRepository {
					m := &mocks.UserRepository{}
					m.On("GetById", mockContext, &id).Return(nil, err).Once()
					return m
				})(),
				hashPassword: func(password string) (string, error) { return password, nil },
				timeService:  &mocks.TimeService{},
			},
			args: args{
				ctx:      mockContext,
				objectID: &id,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test GetById success",
			fields: fields{
				repository: (func(t *testing.T) *mocks.UserRepository {
					m := &mocks.UserRepository{}
					m.On("GetById", mockContext, &id).Return(&user, nil).Once()
					return m
				})(t),
				hashPassword: func(password string) (string, error) { return password, nil },
				timeService:  &mocks.TimeService{},
			},
			args: args{
				ctx:      mockContext,
				objectID: &id,
			},
			want:    &expectedUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceImpl{
				repository:   tt.fields.repository,
				hashPassword: tt.fields.hashPassword,
				timeService:  tt.fields.timeService,
			}
			got, err := s.GetById(tt.args.ctx, tt.args.objectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
			t.Cleanup(func() {
				tt.fields.repository.AssertExpectations(t)
			})
		})
	}
}
