package services

import (
	"encoding/json"
	"log"
	"poc_json/dto"
	"poc_json/models"
	"poc_json/repositories"
	"poc_json/utils"

	"github.com/IBM/sarama"
)

type (
	IDocumentSrvInfs interface {
		GetTest() error
		CreateDocumentVersion(p *models.DocumentVersionModel) error
		GetDocumentVersion(versionID string) (dto.SDocumentVersionResponse, error)
	}
	serviceDeps struct {
		kafkaProducer *sarama.SyncProducer
		documentRepo  repositories.IDocumentRepoInfs
	}
)

func NewDocumentSrv(documentRepo repositories.IDocumentRepoInfs, kafkaProducer *sarama.SyncProducer) IDocumentSrvInfs {
	return &serviceDeps{
		kafkaProducer: kafkaProducer,
		documentRepo:  documentRepo,
	}
}

// CreateDocumentVersion implements IDocumentSrvInfs.
func (s *serviceDeps) CreateDocumentVersion(p *models.DocumentVersionModel) error {
	tx, err := s.documentRepo.BeginTransaction()
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
			return err
		}

		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: models.EVT_NOTIFICATION,
			Value: sarama.StringEncoder(eventJSON),
		})
		if err != nil {
			return err
		}
	}

	if err := s.documentRepo.HelperCreateDocumentVersion(tx, p); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (s *serviceDeps) GetTest() error {
	tx, err := s.documentRepo.BeginTransaction()
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

// GetDocumentVersion implements IDocumentSrvInfs.
func (s *serviceDeps) GetDocumentVersion(versionID string) (dto.SDocumentVersionResponse, error) {
	r, err := s.documentRepo.GetDocumentVersion(versionID)
	if err != nil {
		return dto.SDocumentVersionResponse{}, err
	}
	if r == nil {
		return dto.SDocumentVersionResponse{}, nil
	}

	return dto.SDocumentVersionResponse{
		ID:          r.ID,
		Version:     r.Version,
		VersionType: r.VersionType,
		Value:       r.Value,
	}, nil
}
