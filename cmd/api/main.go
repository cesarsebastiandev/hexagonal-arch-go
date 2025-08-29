package main

import (
	"hexagonal-architecture/cmd/api/handlers/user"
	"hexagonal-architecture/internal/repositories/mongo"
	userMongo "hexagonal-architecture/internal/repositories/mongo/user"
	userService "hexagonal-architecture/internal/services/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()

	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := userMongo.Repository{
		Client: client,
	}

	userSrv := userService.Service{
		Repo: userRepo,
	}

	userHandler := user.Handler{
		UserService: userSrv,
	}

	ginEngine.POST("/api/v1/users", userHandler.CreateUser)

	log.Fatal(ginEngine.Run(":8082"))
}
