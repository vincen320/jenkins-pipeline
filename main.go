package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vincen320/user-service/app"
	"github.com/vincen320/user-service/controller"
	"github.com/vincen320/user-service/helper"
	"github.com/vincen320/user-service/middleware"
	"github.com/vincen320/user-service/repository"
	"github.com/vincen320/user-service/service"
)

func main() {
	/** //Tidak perlu karena di docker-compose bisa set env-file untuk dibaca (pakai ini untuk development aja)
	  		env_file:
	              - .env

	  	err := godotenv.Load()
	  	helper.PanicIfError(err)
	  **/

	//Tidak perlu kalau dideploynya tidak ada file .env
	err := godotenv.Load() //tentu jika ada ENV VAR yang didefinisikan diluar, maka akan menggunakan ENV VAR tsb, bukan dari file .env(memungkinkan kita untuk fleksibel mengganti env var tanpa ganti di file)
	if err != nil {
		log.Println("ERROR LOAD godotenv", err)
	}

	testExecCommandls()

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
		Addr:           ":" + os.Getenv("PORT"), //AGAR BISA DIDEPLOY DI HEROKU
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("User Service Start on" + server.Addr + " port")
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}

func testExecCommandls() {
	//cmd := exec.Command("sh", "-c", "cd 'app' && ls -l -a")
	cmd := exec.Command("sh", "-c", "ls -l -a") // ls -l akan lihat path kita sekarang, bukan path file go ini
	stdout, err := cmd.Output()

	if err != nil {
		log.Println(err.Error())
		return
	}

	// Print the output
	log.Println(string(stdout))
}
