package main

import (
    "consumer-service/kafka"
)

func main() {
    // Start consuming feedback from Kafka
    kafka.ConsumeFeedback()
}
