package kafka

import (
	"consumer-service/db"
	"consumer-service/models"
	"context"
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
