package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var routesMap map[string]interface{}

// Index for the route '/'
func Index(c *gin.Context) {
	routesMap = make(map[string]interface{})
	routesMap["/"] = "[GET] The Root Route"
	routesMap["/feedbacks/create"] = "[POST] Create New Feedback"
	routesMap["/feedbacks"] = "[GET] all feedbacks"
	routesMap["/feedbacks/reply/:feedbackId"] = "[POST] Create New Reply on Feedback with id=feedbackId"

	c.JSON(http.StatusOK,
		gin.H{"status": "OK",
			"message": time.Now().String(),
			"routes":  routesMap,
		})

}
