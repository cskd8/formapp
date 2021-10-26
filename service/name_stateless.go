package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostName struct {
	Name     string `form:"name" validate:"required"`
	Birthday string `form:"birthday"`
	Message  string `form:"message"`
}

func RegisterNameHandler(ctx *gin.Context) {
	var postName PostName
	validate := validator.New()
	if err := ctx.ShouldBind(&postName); err != nil {
		ctx.HTML(http.StatusBadRequest, "name_stateless.html", nil)
		return
	}
	if err := validate.Struct(postName); err != nil {
		ctx.HTML(http.StatusBadRequest, "name_stateless.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "birthday_stateless.html", &postName)
}
