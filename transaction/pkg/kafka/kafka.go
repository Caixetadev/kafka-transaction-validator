package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return &Producer{
		writer: writer,
	}
}

func (p *Producer) SendMessage(ctx context.Context, key string, value interface{}) error {
	var b bytes.Buffer

	if err := json.NewEncoder(&b).Encode(value); err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte(key),
		Value: b.Bytes(),
	}

	err := p.writer.WriteMessages(ctx, message)
	if err != nil {
		log.Fatal("failed to write message:", err)
		return err
	}

	return nil
}

func (p *Producer) Close() {
	p.writer.Close()
}
