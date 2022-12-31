package config

import (
	"context"
	"fmt"

	"github.com/blog-service/src/utils/logger"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(bootstrapServers string) *KafkaProducer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(bootstrapServers),
		AllowAutoTopicCreation: true,
	}
	return &KafkaProducer{writer: w}
}

func (kp KafkaProducer) Produce(ctx context.Context, topic string, message []byte) {
	err := kp.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Value: message,
	})

	if err != nil {
		logger.Error(fmt.Sprintf("delivery failed %s \n", err.Error()), err)
	} else {
		logger.Info(fmt.Sprintf("message delivered topic: %s | key: \n", topic))
	}
}

func (kp KafkaProducer) CloseConnection() {
	kp.writer.Close()
}
