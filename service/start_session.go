package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartSessionHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "name_session.html", nil)
}
