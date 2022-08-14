package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/vincen320/user-service/app"
	"github.com/vincen320/user-service/controller"
	"github.com/vincen320/user-service/helper"
	"github.com/vincen320/user-service/middleware"
	"github.com/vincen320/user-service/repository"
	"github.com/vincen320/user-service/service"
)

func main() {
	/** Tidak perlu karena di docker-compose bisa set env-file untuk dibaca
	  		env_file:
	              - .env

	  	err := godotenv.Load()
	  	helper.PanicIfError(err)
	  **/
	db := app.NewConnection()
	validator := validator.New()
	var UserRepository repository.UserRepository

	DB_DRIVER := os.Getenv("DB_DRIVER")

	if DB_DRIVER == "mysql" {
		UserRepository = repository.NewUserRepository()
	} else {
		UserRepository = repository.NewUserRepositoryPq()
	}

	UserService := service.NewUserService(UserRepository, db, validator)
	UserController := controller.NewUserController(UserService)

	router := gin.New()
	router.Use(middleware.ErrorHandling())
	router.POST("/users", UserController.Create)
	router.PUT("/users/:userId", UserController.Update)
	router.DELETE("/users/:userId", UserController.Delete)
	router.GET("/users/:userId", UserController.FindById)
	router.GET("/users", UserController.FindAll) //gak penting sebenarnya
	router.PATCH("/users/:userId", UserController.UpdatePatch)
	router.GET("/users/username/:username", UserController.FindByUsername)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("User Service Start in 8080 port")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
