package service

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckSessionHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Set("name", "")
	session.Set("birthday", "")
	session.Set("message", "")
	session.Save()
	ctx.HTML(http.StatusOK, "start.html", nil)
}
