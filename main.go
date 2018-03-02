package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.String(200, "Hello World")
}
func setupRouter() (app *gin.Engine) {
	// gin.DisableConsoleColor()
	app = gin.Default()
	// Ping test
	app.GET("/", helloWorld)
	return
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	app := setupRouter()
	fmt.Println("Listening on " + os.Getenv("PORT"))
	// Listen and Server in 0.0.0.0:8080
	app.Run(":" + os.Getenv("PORT"))
}
