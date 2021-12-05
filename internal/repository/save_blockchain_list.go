package repository

import (
	"blockchain_api/pkg/models"
	"fmt"

	"github.com/sirupsen/logrus"
)

func (p *PostgresDB) SaveBlockchainList(blockchainList []models.BlockchainItem) error {

	lenBL := len(blockchainList)
	fmt.Printf("число крипты: %d\n", lenBL)

	_, errExec := p.DB.NamedExec(`INSERT INTO blockchain_list (symbol, price_24h, volume_24h, last_trade_price)
	    VALUES (:symbol, :price_24h, :volume_24h, :last_trade_price)`, blockchainList)
	if errExec != nil {
		logrus.Fatalf("insert error blockchainitem: %s\n", errExec.Error())
	}

	return nil
}
