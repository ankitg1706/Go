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
