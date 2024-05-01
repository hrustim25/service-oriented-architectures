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
	EventID int    `json:"event_id"`
}

type BrokerHandler struct {
	conn *kafka.Conn
}

var statBroker BrokerHandler

func (broker *BrokerHandler) ReadEventMessage() (BrokerMessage, error) {
	// broker.conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	msgBytes, err := broker.conn.ReadMessage(1024)
	if err != nil {
		log.Printf("Read message from broker failed, err: %v", err)
		return BrokerMessage{}, err
	}

	var msg BrokerMessage

	err = json.Unmarshal(msgBytes.Value, &msg)
	if err != nil {
		log.Printf("Json unmarshal error: %v", err)
		return BrokerMessage{}, err
	}

	return msg, nil
}

func SetupAndStartStatMessageBrokerConsumer() {
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

	go func() {
		for {
			msg, err := statBroker.ReadEventMessage()
			if err != nil {
				log.Printf("Read msg error: %v", err)
				continue
			}

			err = statDB.AddEvent(TaskEvent{TaskId: msg.TaskID, UserId: msg.UserID, EventID: msg.EventID})
			if err != nil {
				log.Printf("Write to DB error: %v", err)
				break
			}
		}
	}()
}
