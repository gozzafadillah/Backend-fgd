package users_test

import (
	"charum/business/users"
	_userMock "charum/business/users/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	userRepository _userMock.Repository
	userUseCase    users.UseCase
	userDomain     users.Domain
)

func TestMain(m *testing.M) {
	userUseCase = users.NewUserUseCase(&userRepository)

	userDomain = users.Domain{
		Id:        primitive.NewObjectID(),
		Email:     "test@charum.com",
		Password:  "test123",
		UserName:  "tester",
		Role:      "user",
		IsActive:  true,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	m.Run()
}

func TestUserRegister(t *testing.T) {
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		userRepository.On("GetByEmail", userDomain.Email).Return(users.Domain{}, errors.New("not found")).Once()
		userRepository.On("GetByUsername", userDomain.UserName).Return(users.Domain{}, errors.New("not found")).Once()
		userRepository.On("Create", mock.Anything).Return(userDomain, nil).Once()

		actualUser, token, err := userUseCase.Register(&userDomain)

		assert.NotNil(t, actualUser)
		assert.NotEmpty(t, token)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid Register | Email already registered", func(t *testing.T) {
		expectedErr := errors.New("email is already registered")
		userRepository.On("GetByEmail", userDomain.Email).Return(userDomain, nil).Once()

		actualUser, token, err := userUseCase.Register(&userDomain)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Empty(t, token)
		assert.Equal(t, err, expectedErr)
	})

	t.Run("Test Case 3 | Invalid Register | Username already used", func(t *testing.T) {
		expectedErr := errors.New("username is already used")
		userRepository.On("GetByEmail", userDomain.Email).Return(users.Domain{}, errors.New("not found")).Once()
		userRepository.On("GetByUsername", userDomain.UserName).Return(userDomain, nil).Once()

		actualUser, token, err := userUseCase.Register(&userDomain)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Empty(t, token)
		assert.Equal(t, err, expectedErr)
	})

	t.Run("Test Case 4 | Invalid Register | Error when creating user", func(t *testing.T) {
		expectedErr := errors.New("failed to register user")
		userRepository.On("GetByEmail", userDomain.Email).Return(users.Domain{}, errors.New("not found")).Once()
		userRepository.On("GetByUsername", userDomain.UserName).Return(users.Domain{}, errors.New("not found")).Once()
		userRepository.On("Create", mock.Anything).Return(users.Domain{}, expectedErr).Once()

		actualUser, token, err := userUseCase.Register(&userDomain)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Empty(t, token)
		assert.Equal(t, err, expectedErr)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		copyDomain := userDomain
		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(copyDomain.Password), bcrypt.DefaultCost)
		copyDomain.Password = string(encryptedPassword)

		userRepository.On("GetByEmail", userDomain.Email).Return(copyDomain, nil).Once()

		actualUsers, token, actualErr := userUseCase.Login(&userDomain)

		assert.NotNil(t, actualUsers)
		assert.NotEmpty(t, token)
		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Login | Email not found", func(t *testing.T) {
		expectedErr := errors.New("email is not registered")
		userRepository.On("GetByEmail", userDomain.Email).Return(users.Domain{}, expectedErr).Once()

		actualUsers, token, err := userUseCase.Login(&userDomain)

		assert.Equal(t, users.Domain{}, actualUsers)
		assert.Empty(t, token)
		assert.Equal(t, err, expectedErr)
	})

	t.Run("Test Case 3 | Invalid Login | Wrong password", func(t *testing.T) {
		expectedErr := errors.New("wrong password")
		userRepository.On("GetByEmail", userDomain.Email).Return(userDomain, nil).Once()

		actualUsers, token, actualErr := userUseCase.Login(&userDomain)

		assert.Equal(t, users.Domain{}, actualUsers)
		assert.Empty(t, token)
		assert.Equal(t, actualErr, expectedErr)
	})
}

func TestGetWithSortAndOrder(t *testing.T) {
	t.Run("Test Case 1 | Valid Get Users", func(t *testing.T) {
		userRepository.On("GetWithSortAndOrder", 0, 1, "createdAt", -1).Return([]users.Domain{userDomain}, 1, nil).Once()

		actualUsers, totalData, actualErr := userUseCase.GetWithSortAndOrder(1, 1, "createdAt", "desc")

		assert.NotZero(t, totalData)
		assert.NotNil(t, actualUsers)
		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Get Users | Error when getting users", func(t *testing.T) {
		expectedErr := errors.New("failed to get users")
		userRepository.On("GetWithSortAndOrder", 0, 1, "createdAt", 1).Return([]users.Domain{}, 1, expectedErr).Once()

		actualUsers, totalData, actualErr := userUseCase.GetWithSortAndOrder(1, 1, "createdAt", "asc")

		assert.Zero(t, totalData)
		assert.Equal(t, []users.Domain{}, actualUsers)
		assert.Equal(t, expectedErr, actualErr)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Test Case 1 | Valid Get User", func(t *testing.T) {
		userRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()

		actualUser, actualErr := userUseCase.GetByID(userDomain.Id)

		assert.NotNil(t, actualUser)
		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Get User | Error when getting user", func(t *testing.T) {
		expectedErr := errors.New("failed to get user")
		userRepository.On("GetByID", userDomain.Id).Return(users.Domain{}, expectedErr).Once()

		actualUser, actualErr := userUseCase.GetByID(userDomain.Id)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Valid Update", func(t *testing.T) {
		copyDomain := userDomain
		copyDomain.UserName = "newUsername"
		copyDomain.Email = "newEmail@charum.com"
		copyDomain.DisplayName = "new display"
		copyDomain.IsActive = false

		userRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()
		userRepository.On("GetByEmail", copyDomain.Email).Return(users.Domain{}, errors.New("not found")).Once()
		userRepository.On("GetByUsername", copyDomain.UserName).Return(users.Domain{}, errors.New("not found")).Once()
		userRepository.On("Update", mock.Anything).Return(copyDomain, nil).Once()

		actualUser, actualErr := userUseCase.Update(copyDomain.Id, &copyDomain)

		assert.NotNil(t, actualUser)
		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Update | Error when getting user", func(t *testing.T) {
		expectedErr := errors.New("failed to get user")
		userRepository.On("GetByID", userDomain.Id).Return(users.Domain{}, expectedErr).Once()

		actualUser, actualErr := userUseCase.Update(userDomain.Id, &userDomain)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 3 | Invalid Update | Error when updating user", func(t *testing.T) {
		expectedErr := errors.New("failed to update user")
		userRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()
		userRepository.On("Update", mock.Anything).Return(users.Domain{}, expectedErr).Once()

		actualUser, actualErr := userUseCase.Update(userDomain.Id, &userDomain)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 4 | Invalid Update | Email is already registered", func(t *testing.T) {
		expectedErr := errors.New("email is already registered")
		copyDomain := userDomain
		copyDomain.Email = "existemail@charum.com"
		userRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()
		userRepository.On("GetByEmail", copyDomain.Email).Return(copyDomain, nil).Once()

		actualUser, actualErr := userUseCase.Update(userDomain.Id, &copyDomain)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 5 | Invalid Update | Username is already used", func(t *testing.T) {
		expectedErr := errors.New("username is already used")
		copyDomain := userDomain
		copyDomain.UserName = "existusername"

		userRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()
		userRepository.On("GetByUsername", copyDomain.UserName).Return(copyDomain, nil).Once()

		actualUser, actualErr := userUseCase.Update(userDomain.Id, &copyDomain)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Valid Delete", func(t *testing.T) {
		userRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()
		userRepository.On("Delete", userDomain.Id).Return(nil).Once()

		actualUser, actualErr := userUseCase.Delete(userDomain.Id)

		assert.NotNil(t, actualUser)
		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Delete | Error when getting user", func(t *testing.T) {
		expectedErr := errors.New("failed to get user")
		userRepository.On("GetByID", userDomain.Id).Return(users.Domain{}, expectedErr).Once()

		actualUser, actualErr := userUseCase.Delete(userDomain.Id)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 3 | Invalid Delete | Error when deleting user", func(t *testing.T) {
		expectedErr := errors.New("failed to delete user")
		userRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()
		userRepository.On("Delete", userDomain.Id).Return(expectedErr).Once()

		actualUser, actualErr := userUseCase.Delete(userDomain.Id)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})
}

func TestSuspend(t *testing.T) {
	t.Run("Test Case 1 | Valid Suspend", func(t *testing.T) {
		userRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()
		userRepository.On("Update", mock.Anything).Return(userDomain, nil).Once()

		actualUser, actualErr := userUseCase.Suspend(userDomain.Id)

		assert.NotNil(t, actualUser)
		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Suspend | Error when getting user", func(t *testing.T) {
		expectedErr := errors.New("failed to get user")
		userRepository.On("GetByID", userDomain.Id).Return(users.Domain{}, expectedErr).Once()

		actualUser, actualErr := userUseCase.Suspend(userDomain.Id)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 3 | Invalid Suspend | Error when suspending user", func(t *testing.T) {
		expectedErr := errors.New("failed to suspend user")
		userRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()
		userRepository.On("Update", mock.Anything).Return(users.Domain{}, expectedErr).Once()

		actualUser, actualErr := userUseCase.Suspend(userDomain.Id)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 4 | Invalid Suspend | User is already suspended", func(t *testing.T) {
		expectedErr := errors.New("user is already suspended")
		copyDomain := userDomain
		copyDomain.IsActive = false
		userRepository.On("GetByID", userDomain.Id).Return(copyDomain, nil).Once()

		actualUser, actualErr := userUseCase.Suspend(userDomain.Id)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})
}

func TestUnsuspend(t *testing.T) {
	t.Run("Test Case 1 | Valid Unsuspend", func(t *testing.T) {
		copyDomain := userDomain
		copyDomain.IsActive = false
		userRepository.On("GetByID", userDomain.Id).Return(copyDomain, nil).Once()
		userRepository.On("Update", mock.Anything).Return(userDomain, nil).Once()

		actualUser, actualErr := userUseCase.Unsuspend(userDomain.Id)

		assert.NotNil(t, actualUser)
		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Unsuspend | Error when getting user", func(t *testing.T) {
		expectedErr := errors.New("failed to get user")
		userRepository.On("GetByID", userDomain.Id).Return(users.Domain{}, expectedErr).Once()

		actualUser, actualErr := userUseCase.Unsuspend(userDomain.Id)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 3 | Invalid Unsuspend | Error when unsuspending user", func(t *testing.T) {
		expectedErr := errors.New("failed to unsuspend user")
		copyDomain := userDomain
		copyDomain.IsActive = false
		userRepository.On("GetByID", userDomain.Id).Return(copyDomain, nil).Once()
		userRepository.On("Update", mock.Anything).Return(users.Domain{}, expectedErr).Once()

		actualUser, actualErr := userUseCase.Unsuspend(userDomain.Id)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 4 | Invalid Unsuspend | User is not suspended", func(t *testing.T) {
		expectedErr := errors.New("user is not suspended")
		userRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()

		actualUser, actualErr := userUseCase.Unsuspend(userDomain.Id)

		assert.Equal(t, users.Domain{}, actualUser)
		assert.Equal(t, expectedErr, actualErr)
	})
}
