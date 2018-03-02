package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var routesMap map[string]string

// Index for the route '/'
func Index(c *gin.Context) {
	routesMap = make(map[string]string)
	routesMap["GET"] = "/"
	routesMap["POST"] = "/feedbacks/create"
	routesMap["GET"] = "/feedbacks"
	routesMap["POST"] = "/feedbacks/reply/:feedbackId"

	c.JSON(http.StatusOK,
		gin.H{"status": "OK",
			"message": time.Now().String(),
			"routes":  routesMap,
		})

}
