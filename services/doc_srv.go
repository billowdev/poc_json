package services

import (
	"poc_json/repositories"

	"github.com/IBM/sarama"
)

type (
	IDocumentSrvInfs interface {
		GetTest() string
	}
	documentSrvDeps struct {
		kafkaProducer *sarama.SyncProducer
		repo          repositories.IDocumentRepoInfs
	}
)

func NewDocumentSrv(repo repositories.IDocumentRepoInfs, kafkaProducer *sarama.SyncProducer) IDocumentSrvInfs {
	return &documentSrvDeps{
		kafkaProducer: kafkaProducer,
		repo:          repo,
	}
}

func (s *documentSrvDeps) GetTest() string {
	return s.repo.GetTest()
}
