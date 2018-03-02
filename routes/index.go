package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Index for the route '/'
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": time.Now().String()})

}
