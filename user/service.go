package user

import (
	"context"
	"github.com/eijiok/user-api/dto"
	"github.com/eijiok/user-api/errors"
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/validators"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type serviceImpl struct {
	repository   interfaces.UserRepository
	hashPassword func(password string) (string, error)
}

func NewServiceImpl(repository interfaces.UserRepository, hashPassword func(password string) (string, error)) interfaces.UserService {
	return &serviceImpl{
		repository:   repository,
		hashPassword: hashPassword,
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

func (s *serviceImpl) Save(ctx context.Context, userRequest *dto.UserRequest) (*dto.UserResponse, error) {
	err := validateUser(userRequest, true)
	if err != nil {
		return nil, &errors.ValidationError{
			Errs: []error{err},
		}
	}
	userRequest.Password, err = s.hashPassword(userRequest.Password)
	if err != nil {
		return nil, err
	}
	user := userRequest.ToUser()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	id, err := s.repository.Save(ctx, &user)
	if err != nil {
		return nil, err
	}

	dtoUser := dto.UserResponse{}
	dtoUser.FromUserModel(&user)
	dtoUser.ID = *id

	return &dtoUser, nil
}

func (s *serviceImpl) GetById(ctx context.Context, objectID *primitive.ObjectID) (*dto.UserResponse, error) {
	user, err := s.repository.GetById(ctx, objectID)
	dtoUser := dto.UserResponse{}
	dtoUser.FromUserModel(user)
	return &dtoUser, err
}

func (s *serviceImpl) Update(ctx context.Context, objectID *primitive.ObjectID, userRequest *dto.UserRequest) error {
	err := validateUser(userRequest, false)
	if err != nil {
		return err
	}

	if len(userRequest.Password) > 0 {
		userRequest.Password, err = s.hashPassword(userRequest.Password)
		if err != nil {
			return err
		}
	}

	user := userRequest.ToUser()
	user.ID = *objectID
	user.UpdatedAt = time.Now()

	countUpdated, err := s.repository.Update(ctx, &user)
	log.Printf("updated %d documents", countUpdated)

	return err
}

func (s *serviceImpl) Delete(ctx context.Context, objectId *primitive.ObjectID) error {
	countDeleted, err := s.repository.Delete(ctx, objectId)
	log.Printf("deleted %d documents", countDeleted)
	return err
}

func validateUser(user *dto.UserRequest, required bool) error {
	validationErrors := errors.ValidationError{}
	validationErrors.Append(createNameValidatorFunc(required)(user.Name))
	validationErrors.Append(createEmailValidatorFunc(required)(user.Email))
	validationErrors.Append(createPasswordValidationFunc(required)(user.Password))
	validationErrors.Append(createBirthdayValidatorFunc(false)(user.Birthday))

	if validationErrors.HasErrors() {
		return &validationErrors
	}
	return nil
}

func createNameValidatorFunc(required bool) func(value any) error {
	maxLength := 50
	return newFieldValidator("name", required, validators.ValidateStringLength(nil, &maxLength))
}

func createPasswordValidationFunc(required bool) func(value any) error {
	minLength := 5
	maxLength := 50
	return newFieldValidator("password", required, validators.ValidatorPassword, validators.ValidateStringLength(&minLength, &maxLength))
}

func createEmailValidatorFunc(required bool) func(value any) error {
	return newFieldValidator("email", required, validators.ValidatorEmail)
}

func createBirthdayValidatorFunc(required bool) func(value any) error {
	now := time.Now()
	return newFieldValidator("birthday", required, validators.DateTimeValidator(nil, &now))
}

func newFieldValidator(field string, required bool, validatorSlice ...validators.Validate) func(value any) error {
	return func(value any) error {
		if !required {
			if validators.IsNullOrEmpty(value) {
				return nil
			}
		} else {
			requiredMessage := validators.ValidateRequired(value)
			if len(requiredMessage) > 0 {
				return &errors.ValidationFieldError{
					Field:   field,
					Message: requiredMessage,
				}
			}
		}
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
