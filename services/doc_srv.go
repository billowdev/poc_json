package services

import (
	"encoding/json"
	"log"
	"poc_json/models"
	"poc_json/repositories"
	"poc_json/utils"

	"github.com/IBM/sarama"
)

type (
	IDocumentSrvInfs interface {
		GetTest() error
		CreateDocumentVersion(p models.DocumentVersionModel) error
	}
	serviceDeps struct {
		kafkaProducer *sarama.SyncProducer
		repo          repositories.IDocumentRepoInfs
	}
)

func NewDocumentSrv(repo repositories.IDocumentRepoInfs, kafkaProducer *sarama.SyncProducer) IDocumentSrvInfs {
	return &serviceDeps{
		kafkaProducer: kafkaProducer,
		repo:          repo,
	}
}

// CreateDocumentVersion implements IDocumentSrvInfs.
func (s *serviceDeps) CreateDocumentVersion(p models.DocumentVersionModel) error {
	tx, err := s.repo.BeginTransaction()
	if err != nil {
		tx.Rollback()
		return err
	}
	defer utils.HandleTransaction(tx, err, "CreateDocumentVersion")

	producer := *s.kafkaProducer
	if producer != nil {
		event := models.KafkaEvent{
			Title:   "Version Created",
			Message: "the document version have been created",
		}
		eventJSON, err := json.Marshal(event)
		if err != nil {
			log.Printf("Error marshalling order event: %v", err)
			tx.Rollback()
		}

		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: models.EVT_NOTIFICATION,
			Value: sarama.StringEncoder(eventJSON),
		})
		if err != nil {
			tx.Rollback()
		}
	}
	return nil
}
func (s *serviceDeps) GetTest() error {
	tx, err := s.repo.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("--------------GetTest-----------------")
			err := tx.Rollback()
			log.Println(err)
			log.Println("-------------------------------")
		} else if err != nil {
			tx.Rollback()
		} else {
			if err := tx.Commit().Error; err != nil {
				tx.Rollback()
				log.Println(err)
			}
		}
	}()

	producer := *s.kafkaProducer
	if producer != nil {
		event := models.KafkaEvent{
			Title:   "Order Created",
			Message: "the order have been created",
		}
		eventJSON, err := json.Marshal(event)
		if err != nil {
			log.Printf("Error marshalling order event: %v", err)
			tx.Rollback()
		}

		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: models.EVT_NOTIFICATION,
			Value: sarama.StringEncoder(eventJSON),
		})
		if err != nil {
			tx.Rollback()
		}
	}

	return err
}
