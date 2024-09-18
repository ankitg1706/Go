#!/bin/bash

# Create the main project directory
mkdir -p feedback-system
cd feedback-system || exit

# Create directories for feedback-service
mkdir -p feedback-service/controllers feedback-service/models feedback-service/kafka

# Create directories for consumer-service
mkdir -p consumer-service/kafka consumer-service/db consumer-service/models

# Create feedback-service main.go
cat <<EOL >feedback-service/main.go
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
EOL

# Create feedback-service/controllers/feedback_controller.go
cat <<EOL >feedback-service/controllers/feedback_controller.go
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
EOL

# Create feedback-service/models/feedback.go
cat <<EOL >feedback-service/models/feedback.go
package models

type Feedback struct {
    Username string \`json:"username"\`
    Email    string \`json:"email"\`
    Message  string \`json:"message"\`
}
EOL

# Create feedback-service/kafka/producer.go
cat <<EOL >feedback-service/kafka/producer.go
package kafka

import (
    "context"
    "feedback-service/models"
    "github.com/segmentio/kafka-go"
    "encoding/json"
    "log"
)

func PublishFeedback(feedback models.Feedback) error {
    writer := kafka.Writer{
        Addr:     kafka.TCP("kafka:9092"), // Kafka broker address
        Topic:    "feedback-topic",
        Balancer: &kafka.LeastBytes{},
    }
    defer writer.Close()

    message, err := json.Marshal(feedback)
    if err != nil {
        return err
    }

    err = writer.WriteMessages(context.Background(), kafka.Message{
        Value: message,
    })
    if err != nil {
        log.Println("Failed to write message to Kafka:", err)
        return err
    }

    log.Println("Feedback successfully published to Kafka")
    return nil
}
EOL

# Create consumer-service main.go
cat <<EOL >consumer-service/main.go
package main

import (
    "consumer-service/kafka"
)

func main() {
    // Start consuming feedback from Kafka
    kafka.ConsumeFeedback()
}
EOL

# Create consumer-service/kafka/consumer.go
cat <<EOL >consumer-service/kafka/consumer.go
package kafka

import (
    "context"
    "consumer-service/db"
    "consumer-service/models"
    "encoding/json"
    "log"
    "github.com/segmentio/kafka-go"
)

func ConsumeFeedback() {
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"kafka:9092"},
        Topic:   "feedback-topic",
        GroupID: "feedback-group",
    })

    defer reader.Close()

    for {
        msg, err := reader.ReadMessage(context.Background())
        if err != nil {
            log.Println("Error reading message:", err)
            continue
        }

        var feedback models.Feedback
        err = json.Unmarshal(msg.Value, &feedback)
        if err != nil {
            log.Println("Error unmarshaling message:", err)
            continue
        }

        // Save feedback to database
        err = db.SaveFeedback(feedback)
        if err != nil {
            log.Println("Error saving feedback to database:", err)
        }
    }
}
EOL

# Create consumer-service/db/db.go
cat <<EOL >consumer-service/db/db.go
package db

import (
    "consumer-service/models"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", "postgres://user:password@db:5432/feedback?sslmode=disable")
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("Error pinging the database:", err)
    }

    log.Println("Database connection established")
}

func SaveFeedback(feedback models.Feedback) error {
    query := \`INSERT INTO feedback (username, email, message) VALUES (\$1, \$2, \$3)\`
    _, err := db.Exec(query, feedback.Username, feedback.Email, feedback.Message)
    if err != nil {
        return err
    }

    log.Println("Feedback saved to the database")
    return nil
}
EOL

# Create consumer-service/models/feedback.go
cat <<EOL >consumer-service/models/feedback.go
package models

type Feedback struct {
    Username string
    Email    string
    Message  string
}
EOL

# Create Dockerfile for feedback-service
cat <<EOL >feedback-service/Dockerfile
FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o feedback-service .

CMD ["./feedback-service"]
EOL

# Create Dockerfile for consumer-service
cat <<EOL >consumer-service/Dockerfile
FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o consumer-service .

CMD ["./consumer-service"]
EOL

# Create docker-compose.yml
cat <<EOL >docker-compose.yml
version: '3.8'

services:
  zookeeper:
    image: wurstmeister/zookeeper:3.4.6
    ports:
      - "2181:2181"

  kafka:
    image: wurstmeister/kafka:2.12-2.2.1
    ports:
      - "9092:9092"
    environment:
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka:9092
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
    depends_on:
      - zookeeper

  db:
    image: postgres:13
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: feedback
    ports:
      - "5432:5432"
    volumes:
      - postgres-data:/var/lib/postgresql/data

  feedback-service:
    build: ./feedback-service
    ports:
      - "8080:8080"
    depends_on:
      - kafka

  consumer-service:
    build: ./consumer-service
    depends_on:
      - db
      - kafka

volumes:
  postgres-data:
EOL

# Print success message
echo "Project structure created successfully!"

