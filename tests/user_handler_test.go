package tests

import (
	"net/http/httptest"
	"testing"
	"userapi/domain"
	"userapi/mocks"
	"userapi/usecases"

	handler "userapi/adapters/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByIdHandler(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	app := fiber.New()

	//define the test input and output
	userId := 1
	execptedUser := &domain.User{
		ID:       1,
		Username: "John Doe",
		Email:    "john.doe@gmail.com",
	}
	//set up the mock behavior
	mockRepo.On("GetUserByID", userId).Return(execptedUser, nil)

	//Initalize the use case and handler

	userUC := &usecases.UserUsecase{UserRepository: mockRepo}
	userHandler := handler.UserHandler{UseCase: userUC}

	//set up the route
	app.Get("/user/:id", userHandler.GetUserByID)

	//createt the request
	req := httptest.NewRequest("GET", "/user/1", nil)
	resp, err := app.Test(req)

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	mockRepo.AssertExpectations(t)
}
