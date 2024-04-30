package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	BROKER_TOPIC_ENV = "BROKER_TOPIC_ENV"
)

const (
	ViewEventID = iota
	LikeEventID
)

type BrokerMessage struct {
	UserID  uint64 `json:"user_id"`
	TaskID  uint64 `json:"task_id"`
	EventID uint64 `json:"event_id"`
}

type BrokerHandler struct {
	conn *kafka.Conn
}

var statBroker BrokerHandler

func (broker *BrokerHandler) SendEventMessage(msg BrokerMessage) error {
	buf, err := json.Marshal(&msg)
	if err != nil {
		return err
	}

	// broker.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	wrote := 0
	for wrote < len(buf) {
		val, err := broker.conn.Write(buf[wrote:])
		if err != nil {
			log.Printf("Error while writing to broker: %v", err)
		}

		wrote += val
	}
	return nil
}

func SetupStatMessageBroker() {
	if _, ok := os.LookupEnv(BROKER_TOPIC_ENV); !ok {
		log.Fatalf("'%v' env var not found", BROKER_TOPIC_ENV)
	}

	topicName := os.Getenv(BROKER_TOPIC_ENV)
	partition := 0

	time.Sleep(5 * time.Second)

	conn, err := kafka.DialLeader(context.Background(), "tcp", "stat_broker:9092", topicName, partition)
	if err != nil {
		log.Fatalf("Failed to connect to broker, err: %v", err)
	}

	statBroker = BrokerHandler{conn}
	_ = statBroker
}
