package config

import (
	"os"
)

type KafkaConfig struct {
	Brokers     []string
	GroupID     string
	Topics      []string
	Concurrency int
}

func NewKafkaConfig() KafkaConfig {
	// Customize your Kafka configuration here
	return KafkaConfig{
		Brokers:     []string{GetKafkaURL()},
		GroupID:     "hrm-consumer-group",
		Topics:      []string{"user-created", "leave-events"},
		Concurrency: 2,
	}
}

func GetKafkaURL() string {
	return os.Getenv("KAFKA_HOST")
}
