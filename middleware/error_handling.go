package middleware

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vincen320/user-service/exception"
	"github.com/vincen320/user-service/model/web"
)

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				log.Println(reflect.TypeOf(err)) //debugging
				switch err.(type) {
				case exception.BadRequestErr:
					BadRequestErrResponse(c, err)
				case exception.NotFoundErr:
					NotFoundErrResponse(c, err)
				case validator.ValidationErrors:
					ValidationErrorResponse(c, err)
				default:
					c.JSON(http.StatusInternalServerError, web.WebResponse{
						Status:  http.StatusInternalServerError,
						Message: "Internal Server Error",
						Data:    nil,
					})
				}
			}
		}()

		c.Next()
	}
}

func BadRequestErrResponse(c *gin.Context, err interface{}) {
	badRequest, _ := err.(exception.BadRequestErr)
	c.JSON(http.StatusBadRequest, web.WebResponse{
		Status:  http.StatusBadRequest,
		Message: badRequest.Error(),
		Data:    nil,
	})
}

func NotFoundErrResponse(c *gin.Context, err interface{}) {
	notFound, _ := err.(exception.NotFoundErr)
	c.JSON(http.StatusNotFound, web.WebResponse{
		Status:  http.StatusNotFound,
		Message: notFound.Error(),
		Data:    nil,
	})
}

func ValidationErrorResponse(c *gin.Context, err interface{}) {
	validationErr, _ := err.(validator.ValidationErrors)
	c.JSON(http.StatusBadRequest, web.WebResponse{
		Status:  http.StatusBadRequest,
		Message: validationErr.Error(),
		Data:    nil,
	})
}
