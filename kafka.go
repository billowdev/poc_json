package main

import (
	"log"

	"github.com/IBM/sarama"
)

func initSaramaProducer() *sarama.SyncProducer {
	// Initialize Sarama configuration
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	brokerURLs := KAFKA_BROKER_URLS
	// Initialize Sarama sync producer
	producer, err := sarama.NewSyncProducer(brokerURLs, config)
	if err != nil {
		log.Fatalf("Error creating Sarama producer: %v", err)
	}

	return &producer
}

func initSaramaConsumer() *sarama.ConsumerGroup {
	// Initialize Sarama configuration
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	brokerURLs := KAFKA_BROKER_URLS
	kafkaConsumerGroup := KAFKA_CONSUMER_GROUP
	// Initialize Kafka consumer group
	consumerGroup, err := sarama.NewConsumerGroup(brokerURLs, kafkaConsumerGroup, config)
	if err != nil {
		log.Fatalf("Error creating Sarama consumer group: %v", err)
	}

	return &consumerGroup
}
