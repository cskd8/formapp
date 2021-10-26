package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "start.html", nil)
}
