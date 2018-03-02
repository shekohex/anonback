package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shekohex/anonback/routes"
)

func setupRouter() (app *gin.Engine) {
	// gin.DisableConsoleColor()
	app = gin.Default()
	app.GET("/", routes.Index)
	app.GET("/feedbacks", routes.GetFeedbacks)
	app.POST("/feedbacks/reply/:feedbackId", routes.ReplyFeedback)
	app.POST("/feedbacks/create", routes.CreateFeedback)
	return
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	app := setupRouter()
	fmt.Println("Listening on " + os.Getenv("PORT"))
	// Listen and Server in 0.0.0.0:8080
	app.Run(":" + os.Getenv("PORT"))
}
