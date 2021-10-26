package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostBirthday struct {
	Name     string `form:"name" validate:"required,min=1"`
	Birthday string `form:"birthday" validate:"required"`
	Message  string `form:"message"`
}

func RegisterBirthdayHandler(ctx *gin.Context) {
	var birthday PostBirthday
	validate := validator.New()
	if err := ctx.ShouldBind(&birthday); err != nil {
		ctx.HTML(http.StatusBadRequest, "birthday_stateless.html", gin.H{"Name": birthday.Name})
		return
	}
	if err := validate.Struct(&birthday); err != nil {
		ctx.HTML(http.StatusBadRequest, "birthday_stateless.html", gin.H{"Name": birthday.Name})
		return
	}
	ctx.HTML(http.StatusOK, "message_stateless.html", &birthday)
}
