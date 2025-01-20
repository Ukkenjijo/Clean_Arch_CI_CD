package tests

import (
	"testing"
	"userapi/domain"
	"userapi/mocks"
	"userapi/usecases"

	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUsecase := usecases.UserUsecase{UserRepository: mockRepo}

	excptedUser := &domain.User{
		ID:    1,
		Username:  "John Doe",
		Email: "john.doe@example.com",
	}
	mockRepo.On("GetUserByID", 1).Return(excptedUser, nil)

	//Act
	user, err := userUsecase.GetUserByID(1)

	//Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t,excptedUser.Username,user.Username)
	mockRepo.AssertCalled(t,"GetUserByID",1)
}
