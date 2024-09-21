package main

import (
    "github.com/gin-gonic/gin"
    "feedback-service/controllers"
)

func main() {
    router := gin.Default()

    // Route to handle feedback submission
    router.POST("/feedback", controllers.SubmitFeedback)

    router.Run(":8080")
}
