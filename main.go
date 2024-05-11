package main

import (
	"alexnerotd/apps/crud_user/handlers"
	"alexnerotd/apps/crud_user/models"
	"alexnerotd/apps/crud_user/repositories"
	"alexnerotd/apps/crud_user/services"
	"alexnerotd/apps/crud_user/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := utils.ConnectionDB()
	if err != nil {
		panic("error connection to database")
	}

	user := models.User{Name: "Jesus", Email: "jesalemalo@gmail.com"}

	result := db.Create(&user)
	fmt.Println(result)
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("error migrating User table")
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewHandlerUser(userService)

	router := gin.Default()

	router.GET("/users/:id", userHandler.GetUserById)
	router.GET("/", userHandler.GetUserById)

	router.Run(":8080")

}
