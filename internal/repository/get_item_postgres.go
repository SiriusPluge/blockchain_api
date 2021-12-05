package repository

import (
	"blockchain_api/pkg/models"
	"fmt"

	"github.com/sirupsen/logrus"
)

type GetItemPostgres struct {
	db *PostgresDB
}

func NewGetItemPostgres(db *PostgresDB) *GetItemPostgres {
	return &GetItemPostgres{db: db}
}

func (s *GetItemPostgres) GetBlockchainItem(cryptName string) (models.BlockchainItemPostgres, error) {

	var item models.BlockchainItemPostgres

	query := fmt.Sprintf("SELECT * FROM blockchain_list WHERE symbol = '%s'", cryptName)
	err := s.db.DB.Get(&item, query)
	if err != nil {
		logrus.Fatalf("get error blockchain item in DB: %s", err.Error())
		return item, err
	}

	return item, nil
}
