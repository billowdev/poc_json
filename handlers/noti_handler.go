package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"poc_json/models"
	"poc_json/services"

	"github.com/IBM/sarama"
)

type (
	INotiHandlerInfs interface {
		StartConsuming(topics []string)
		sarama.ConsumerGroupHandler
		Cleanup(sarama.ConsumerGroupSession) error
		Setup(sarama.ConsumerGroupSession) error
	}
	notiHandlerDeps struct {
		notiSrv       services.INotiSrvInfs
		kafkaConsumer *sarama.ConsumerGroup
	}
)

func NewNotiHandler(notiSrv services.INotiSrvInfs, kafkaConsumer *sarama.ConsumerGroup) INotiHandlerInfs {
	return &notiHandlerDeps{
		notiSrv:       notiSrv,
		kafkaConsumer: kafkaConsumer,
	}
}

// Cleanup implements INotiHandlerInfs.
func (n *notiHandlerDeps) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// Setup implements INotiHandlerInfs.
func (n *notiHandlerDeps) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// StartConsuming implements INotiHandlerInfs.
func (n *notiHandlerDeps) StartConsuming(topics []string) {
	go func() {
		for {
			if n.kafkaConsumer != nil {
				consumer := *n.kafkaConsumer
				err := consumer.Consume(context.Background(), topics, n)
				if err != nil {
					log.Printf("Error consuming Kafka message: %v", err)
				}
			}
		}
	}()
}

func (h *notiHandlerDeps) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("-------------notificationHandlerDeps-------------------")
			err := fmt.Errorf("panic occurred ConsumeClaim: %v", r)
			log.Println(err)
			log.Println("--------------------------------")
		}
	}()
	for kafkaMessage := range claim.Messages() {
		// var event map[string]interface{}
		// if err := json.Unmarshal(kafkaMessage.Value, &event); err != nil {
		// 	log.Printf("Error unmarshalling Kafka message: %v", err)
		// 	continue
		// }
		var kafkaEvent models.KafkaEvent
		if err := json.Unmarshal(kafkaMessage.Value, &kafkaEvent); err != nil {
			log.Printf("Error unmarshalling Kafka message: %v", err)
			continue
		}

		notification := &models.NotificationModel{
			Title:   kafkaEvent.Title,
			Message: kafkaEvent.Message,
		}

		// Save the notification to the database
		if err := h.notiSrv.CreateNotification(notification); err != nil {
			log.Printf("Error creating notification: %v", err)
			continue
		}
		session.MarkMessage(kafkaMessage, "")

	}

	return nil
}
