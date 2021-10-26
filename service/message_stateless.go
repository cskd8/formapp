package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostMessage struct {
	Name     string `form:"name" validate:"required,min=1"`
	Birthday string `form:"birthday" validate:"required"`
	Message  string `form:"message" validate:"required"`
}

func RegisterMessageHandler(ctx *gin.Context) {
	var postMessage PostMessage
	validate := validator.New()
	if err := ctx.ShouldBind(&postMessage); err != nil {
		ctx.HTML(http.StatusBadRequest, "message_stateless.html", gin.H{"Name": postMessage.Name, "Birthday": postMessage.Birthday})
		return
	}
	if err := validate.Struct(postMessage); err != nil {
		ctx.HTML(http.StatusBadRequest, "message_stateless.html", gin.H{"Name": postMessage.Name, "Birthday": postMessage.Birthday})
		return
	}
	ctx.HTML(http.StatusOK, "check_stateless.html", &postMessage)
}
