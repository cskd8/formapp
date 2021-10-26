package service

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostMessageSession struct {
	Message string `form:"message" validate:"required"`
}

func RegisterMessageSessionHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var postMessage PostMessageSession
	validate := validator.New()
	if err := ctx.ShouldBind(&postMessage); err != nil {
		ctx.HTML(http.StatusBadRequest, "message_session.html", nil)
		return
	}
	if err := validate.Struct(postMessage); err != nil {
		message, ok := session.Get("message").(string)
		if !ok {
			message = ""
		}
		ctx.HTML(http.StatusOK, "message_session.html", gin.H{"Message": message})
		return
	}
	session.Set("message", postMessage.Message)
	session.Save()
	name, ok := session.Get("name").(string)
	if !ok {
		name = ""
	}
	birthday, ok := session.Get("birthday").(string)
	if !ok {
		birthday = ""
	}
	ctx.HTML(http.StatusOK, "check_session.html", gin.H{"Name": name, "Birthday": birthday, "Message": postMessage.Message})
}
