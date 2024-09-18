package controllers

import (
    "feedback-service/kafka"
    "feedback-service/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func SubmitFeedback(c *gin.Context) {
    var feedback models.Feedback

    // Bind JSON request to feedback model
    if err := c.BindJSON(&feedback); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Send feedback to Kafka
    err := kafka.PublishFeedback(feedback)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process feedback"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "Feedback submitted successfully"})
}
