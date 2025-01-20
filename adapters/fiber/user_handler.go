package fiber

import (
	"strconv"
	"userapi/domain"
	"userapi/usecases"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UseCase *usecases.UserUsecase
}

func (h *UserHandler) CreateUser(f *fiber.Ctx) error {

	var user domain.User
	if err := f.BodyParser(&user); err != nil {
		return f.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.UseCase.CreateUser(&user); err != nil {
		return f.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})

}

func (h *UserHandler) GetUserByID(f *fiber.Ctx) error {
	id, err := strconv.Atoi(f.Params("id"))
	if err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user,err:=h.UseCase.GetUserByID(id)
	if err!=nil {
		return f.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return f.Status(fiber.StatusOK).JSON(user)

}
func (h *UserHandler) UpdateUser(f *fiber.Ctx) error {
	var user domain.User
	if err := f.BodyParser(&user); err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
    if err:=h.UseCase.UpdateUser(&user);err!=nil{
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
	

}
func (h *UserHandler) DeleteUser(f *fiber.Ctx) error {
	id, err := strconv.Atoi(f.Params("id"))
	if err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err:=h.UseCase.DeleteUser(id);err!=nil{
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
