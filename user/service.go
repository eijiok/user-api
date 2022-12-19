package user

import (
	"context"
	"github.com/eijiok/user-api/dto"
	"github.com/eijiok/user-api/errors"
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/model"
	"github.com/eijiok/user-api/validators"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type serviceImpl struct {
	repository interfaces.UserRepository
}

func NewServiceImpl(repository interfaces.UserRepository) interfaces.UserService {
	return &serviceImpl{
		repository: repository,
	}
}

func (s *serviceImpl) List(ctx context.Context) ([]dto.UserResponse, error) {

	list, err := s.repository.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]dto.UserResponse, len(list))
	for i, user := range list {
		dtoUser := dto.UserResponse{}
		dtoUser.FromUserModel(&user)
		result[i] = dtoUser
	}

	return result, err
}

func (s *serviceImpl) Save(ctx context.Context, user *model.User) (*dto.UserResponse, error) {
	err := validateUser(user)
	if err != nil {
		return nil, err
	}
	err = validateRequiredPassword()(user.Password)
	if err != nil {
		return nil, &errors.ValidationError{
			Errs: []error{err},
		}
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	id, err := s.repository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	dtoUser := dto.UserResponse{}
	dtoUser.FromUserModel(user)
	dtoUser.ID = *id

	return &dtoUser, nil
}

func validateUser(user *model.User) error {
	validationErrors := errors.ValidationError{}
	validationErrors.Append(nameValidator()(user.Name))
	validationErrors.Append(validateBirthday()(user.Birthday))
	validationErrors.Append(validateEmail()(user.Email))

	if validationErrors.HasErrors() {
		return &validationErrors
	}
	return nil
}

func newFieldValidator(field string, validatorSlice ...validators.Validate) func(value any) error {
	return func(value any) error {
		for _, validator := range validatorSlice {
			errorMessage := validator(value)
			if len(errorMessage) > 0 {
				return &errors.ValidationFieldError{
					Field:   field,
					Message: errorMessage,
				}
			}
		}
		return nil
	}
}

func nameValidator() func(value any) error {
	maxLength := 50
	return newFieldValidator("name", validators.ValidateRequired, validators.ValidateStringLength(nil, &maxLength))
}

func validateRequiredPassword() func(value any) error {
	return newFieldValidator("password", validators.ValidateRequired, validators.ValidatorPassword)
}
func validatePassword() func(value any) error {
	return newFieldValidator("password", validators.ValidatorPassword)
}

func validateEmail() func(value any) error {
	return newFieldValidator("email", validators.ValidateRequired, validators.ValidatorEmail)
}

func validateBirthday() func(value any) error {
	now := time.Now()
	return newFieldValidator("birthday", validators.DateTimeValidator(nil, &now))
}

func (s *serviceImpl) GetById(ctx context.Context, objectID *primitive.ObjectID) (*dto.UserResponse, error) {
	user, err := s.repository.GetById(ctx, objectID)
	dtoUser := dto.UserResponse{}
	dtoUser.FromUserModel(user)
	return &dtoUser, err
}

func (s *serviceImpl) Update(ctx context.Context, objectID *primitive.ObjectID, user *model.User) error {
	user.ID = *objectID
	user.UpdatedAt = time.Now()
	err := validateUser(user)
	if err != nil {
		return err
	}
	err = validatePassword()(user.Password)
	if err != nil {
		return &errors.ValidationError{
			Errs: []error{err},
		}
	}
	countUpdated, err := s.repository.Update(ctx, user)
	log.Printf("updated %d documents", countUpdated)

	return err
}

func (s *serviceImpl) Delete(ctx context.Context, objectId *primitive.ObjectID) error {
	countDeleted, err := s.repository.Delete(ctx, objectId)
	log.Printf("deleted %d documents", countDeleted)
	return err
}
