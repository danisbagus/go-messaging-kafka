package modules

import (
	"context"
	"os"

	"github.com/segmentio/kafka-go"
)

type ConsumeHandler func(ctx context.Context, message kafka.Message) error

type KafkaConsumer struct {
	Reader *kafka.Reader
}

func NewKafkaConsumer() *KafkaConsumer {
	brokers := []string{
		os.Getenv("KAFKA_BROKER_1"),
	}

	groupTopics := []string{
		os.Getenv("TOPIC_TRANSACTION"),
	}

	groupID := os.Getenv("GROUP_ID_TRANSACTION")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		GroupID:     groupID,
		GroupTopics: groupTopics,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
	})

	return &KafkaConsumer{Reader: reader}
}
