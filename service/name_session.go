package service

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostNameSession struct {
	Name string `form:"name" validate:"required"`
}

func RegisterNameSessionHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var postData PostNameSession
	validate := validator.New()
	if err := ctx.ShouldBind(&postData); err != nil {
		ctx.HTML(http.StatusBadRequest, "name_session.html", nil)
		return
	}
	if err := validate.Struct(&postData); err != nil {
		name, ok := session.Get("name").(string)
		if !ok {
			name = ""
		}
		ctx.HTML(http.StatusOK, "name_session.html", gin.H{"Name": name})
		return
	}
	session.Set("name", postData.Name)
	session.Save()
	birthday, ok := session.Get("birthday").(string)
	if !ok {
		birthday = ""
	}
	ctx.HTML(http.StatusOK, "birthday_session.html", gin.H{"Birthday": birthday})
}
