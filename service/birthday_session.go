package service

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostBirthdaySession struct {
	Birthday string `form:"birthday" validate:"required"`
}

func RegisterBirthdaySessionHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var postBirthdaySession PostBirthdaySession
	validate := validator.New()
	if err := ctx.ShouldBind(&postBirthdaySession); err != nil {
		ctx.HTML(http.StatusBadRequest, "birthday_session.html", nil)
		return
	}
	if err := validate.Struct(&postBirthdaySession); err != nil {
		birthday, ok := session.Get("birthday").(string)
		if !ok {
			birthday = ""
		}
		ctx.HTML(http.StatusOK, "birthday_session.html", gin.H{"Birthday": birthday})
		return
	}
	session.Set("birthday", postBirthdaySession.Birthday)
	session.Save()

	message, ok := session.Get("message").(string)
	if !ok {
		message = ""
	}
	ctx.HTML(http.StatusOK, "message_session.html", gin.H{"Message": message})
}
