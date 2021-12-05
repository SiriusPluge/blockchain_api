package service

import (
	"blockchain_api/internal/repository"
	"blockchain_api/pkg/models"
)

type Blockchain interface {
	GetBlockchainItem(cryptName string) (models.BlockchainItemPostgres, error)
}

type Service struct {
	Blockchain
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Blockchain: NewGetService(repos.Blockchain),
	}
}
