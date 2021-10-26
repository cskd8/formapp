package main

import (
	"fmt"
	"net/http"

	"formapp.go/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Text Email Number Range Checkbox
type InputData struct {
	Text     string `form:"text"`
	Email    string `form:"email"`
	Number   string `form:"number"`
	Range    string `form:"range"`
	Checkbox string `form:"checkbox"`
}

// config
const port = 8000

func main() {
	// initialize Gin engine
	engine := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("mysession", store))

	// routing
	engine.GET("/", rootHandler)
	engine.POST("/name", service.StartHandler)
	engine.POST("/name_session", service.StartSessionHandler)
	engine.POST("/start", service.CheckHandler)
	engine.POST("/start_session", service.CheckSessionHandler)
	engine.POST("/birthday", service.RegisterNameHandler)
	engine.POST("/birthday_session", service.RegisterNameSessionHandler)
	engine.POST("/message", service.RegisterBirthdayHandler)
	engine.POST("/message_session", service.RegisterBirthdaySessionHandler)
	engine.POST("/check", service.RegisterMessageHandler)
	engine.POST("/check_session", service.RegisterMessageSessionHandler)

	engine.LoadHTMLGlob("templates/*.html")

	// start server
	engine.Run(fmt.Sprintf(":%d", port))
}

func rootHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "start.html", nil)
}
