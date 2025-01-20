package infrastructure

import (
	"userapi/adapters/fiber"
	f "github.com/gofiber/fiber/v2")

func NewRouter(UserHandler *fiber.UserHandler) *f.App {
	app := f.New()
	app.Post("/user", UserHandler.CreateUser)
	app.Get("/user/:id", UserHandler.GetUserByID)
	app.Put("/user", UserHandler.UpdateUser)
	app.Delete("/user/:id", UserHandler.DeleteUser)
	return app
}
