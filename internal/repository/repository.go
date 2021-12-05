package repository

import (
	"blockchain_api/pkg/models"
)

type Blockchain interface {
	GetBlockchainItem(cryptName string) (models.BlockchainItemPostgres, error)
}

type Repository struct {
	Blockchain
}

func NewRepositiry(db *PostgresDB) *Repository {
	return &Repository{
		Blockchain: NewGetItemPostgres(db),
	}
}
