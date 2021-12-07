package service

import (
	"blockchain_api/internal/repository"
	"blockchain_api/pkg/models"
)

type GetService struct {
	repo repository.Blockchain
}

func NewGetService(repo repository.Blockchain) *GetService {
	return &GetService{repo: repo}
}

func (g *GetService) GetBlockchainItem(cryptName string) (models.BlockchainItemPostgres, error) {
	return g.repo.GetBlockchainItem(cryptName)
}
