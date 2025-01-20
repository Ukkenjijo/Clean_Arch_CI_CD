package main

import (
	"context"
	"log"
	"userapi/adapters/db"
	"userapi/adapters/fiber"
	"userapi/infrastructure"
	"userapi/usecases"
)

func main() {

	client,err:=db.NewMongoClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	userRepo := db.NewMongoUserRepository(client, "userdb", "users")
	UserUsecase := &usecases.UserUsecase{UserRepository: userRepo}
	UserHandler := &fiber.UserHandler{UseCase: UserUsecase}
	router := infrastructure.NewRouter(UserHandler)
	//start the server
	log.Fatal(router.Listen(":3000"))
}
