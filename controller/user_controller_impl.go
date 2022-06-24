package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincen320/user-service/exception"
	"github.com/vincen320/user-service/model/web"
	"github.com/vincen320/user-service/service"
)

//us *UserControllerImpl
//pada baris diatas tekan CTRL +SHIFT+ P
//Pilih Go Generate Interface Stubs
//ketik package.Interface >> controller.UserController

type UserControllerImpl struct {
	Service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		Service: service,
	}
}

func (uc *UserControllerImpl) Create(c *gin.Context) {
	var createUserRequest web.UserCreateRequest
	err := c.ShouldBind(&createUserRequest)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewBadRequestErr(err.Error())) //Bad Request
	} //SUATU SAAT MUNGKIN BISA BUAT FUNGSI ERROR SENDIRI DI EXCEPTION ?, BADREQUESTREPONSEIFERROR(ERR)
	result := uc.Service.Create(c, createUserRequest)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Successfully Create User",
		Data:    result,
	})
}

func (uc *UserControllerImpl) Update(c *gin.Context) {
	var updateUserRequest web.UserUpdateRequest
	err := c.ShouldBind(&updateUserRequest)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewBadRequestErr(err.Error())) //Bad Request
	}
	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewBadRequestErr(err.Error())) //Bad Request
	}
	updateUserRequest.Id = id

	result := uc.Service.Update(c, updateUserRequest)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Successfully Update User",
		Data:    result,
	})
}

func (uc *UserControllerImpl) Delete(c *gin.Context) {
	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewBadRequestErr(err.Error())) //Bad Request
	}
	result := uc.Service.Delete(c, id)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Successfully Delete User",
		Data:    result,
	})
}

func (uc *UserControllerImpl) FindById(c *gin.Context) {
	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewBadRequestErr(err.Error())) //Bad Request
	}
	result := uc.Service.FindById(c, id)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Find User with id" + userId,
		Data:    result,
	})
}

func (uc *UserControllerImpl) FindAll(c *gin.Context) {
	result := uc.Service.FindAll(c)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Succesfully find all user",
		Data:    result,
	})
}

func (uc *UserControllerImpl) UpdatePatch(c *gin.Context) {
	var updatePatchRequest web.UserUpdatePatchRequest
	err := c.ShouldBind(&updatePatchRequest)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewBadRequestErr(err.Error())) // Bad Request
	}

	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil { ////harus begini kalau yang pakai customError karena kalau pakai helper langsung, bisa saja err nya itu tidak ada, sehingga kalau dipaksakan dengan helper akan error karena errnya berarti != nil karena parameternya ada nilai yaitu structnya
		panic(exception.NewBadRequestErr(err.Error())) //Bad Request
	}
	updatePatchRequest.Id = id
	result := uc.Service.UpdatePatch(c, updatePatchRequest)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Successfully Update User",
		Data:    result,
	})
}

func (uc *UserControllerImpl) FindByUsername(c *gin.Context) {
	username := c.Param("username")

	result := uc.Service.FindByUsername(c, username)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Find User with username " + username,
		Data:    result,
	})
}
