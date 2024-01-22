package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topic,
		Partition: 10,
		MaxBytes:  10e6,
	})

	reader.SetOffset(0)

	return &Consumer{reader: reader}
}

func (c *Consumer) ReadMessage(ctx context.Context, value interface{}) {
	m, err := c.reader.ReadMessage(context.Background())
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(m.Value, value)
	if err != nil {
		log.Println(err)
	}
}

func (c *Consumer) Close() {
	c.reader.Close()
}
