package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shekohex/anonback/config"
)

// CreateFeedback create New Feedback
func CreateFeedback(c *gin.Context) {
	var msg config.Message
	err := c.BindJSON(&msg)
	if err == nil {
		msg.CreatedAt = time.Now().String()
		feedback := new(config.Feedback)
		feedback.New(&msg)
		go func() {
			bin := feedback.MarshalBinary()
			f, err := config.Redis.HSet("feedbacks", feedback.ID, bin).Result()
			if err == nil {
				fmt.Printf("Feedback #%s Saved: %v\n", feedback.ID, f)
			} else {
				fmt.Printf("Error While Saving Feedback: %s\n", err)
			}
		}()
		c.JSON(http.StatusCreated, gin.H{"status": "Created", "message": &feedback})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest", "message": err})
	}
}

// GetFeedbacks get all feedbacks
func GetFeedbacks(c *gin.Context) {
	feedbacks, err := config.Redis.HVals("feedbacks").Result()
	var newFeedback config.Feedback
	var data [][]byte
	for _, element := range feedbacks {
		json.Unmarshal([]byte(element), &newFeedback)
		val, _ := json.Marshal(newFeedback)
		data = append(data, val)
	}
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": "OK", "message": data})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError", "message": err})
	}
}

// ReplyFeedback reply for one of the feedbacks
func ReplyFeedback(c *gin.Context) {
	var msg config.Message
	err := c.BindJSON(&msg)
	if err == nil {
		msg.CreatedAt = time.Now().String()
		var feedback config.Feedback
		go func() {
			data, err := config.Redis.HGet("feedbacks", c.Param("feedbackId")).Result()
			json.Unmarshal([]byte(data), &feedback)
			feedback.Reply = &msg
			bin, _ := json.Marshal(&feedback)
			f, err := config.Redis.HSet("feedbacks", feedback.ID, bin).Result()
			if err == nil {
				fmt.Printf("Feedback #%s Saved: %+v\n", feedback.ID, !f)
			} else {
				fmt.Printf("Error While Saving Feedback: %s\n", err)
			}
		}()
		c.JSON(http.StatusCreated, gin.H{"status": "Created", "message": "OK"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest", "message": err})
	}

}
