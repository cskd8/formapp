package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostStart struct {
	Name     string `form:"name"`
	Birthday string `form:"birthday"`
	Message  string `form:"message"`
}

func StartHandler(ctx *gin.Context) {
	var post PostStart
	if err := ctx.ShouldBind(&post); err != nil {
		ctx.HTML(http.StatusBadRequest, "start.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "name_stateless.html", &post)
}
