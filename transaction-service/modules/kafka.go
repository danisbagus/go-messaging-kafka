package modules

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
	Dialer *kafka.Dialer
}

func CreateTopics(topics []string) {
	if len(topics) == 0 {
		fmt.Println("topics cannot empty")

	}
	connection, err := kafka.Dial("tcp", net.JoinHostPort("localhost", "9092"))
	if err != nil {
		panic(err.Error())
	}

	topicConfigs := make([]kafka.TopicConfig, 0)

	for _, topic := range topics {
		topicConfig := kafka.TopicConfig{Topic: topic, NumPartitions: 1, ReplicationFactor: 1}
		topicConfigs = append(topicConfigs, topicConfig)
	}

	err = connection.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

func NewKafkaProducer() *KafkaProducer {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{
			os.Getenv("KAFKA_BROKER_1"),
		},
		Dialer: dialer,
	})

	return &KafkaProducer{
		Writer: writer,
	}
}

func (kp *KafkaProducer) Produce(key []byte, value []byte, topic string) error {
	err := kp.Writer.WriteMessages(context.TODO(), kafka.Message{
		Topic:  topic,
		Offset: 0,
		Key:    key,
		Value:  value,
	})

	if err != nil {
		return err
	}

	return nil
}
