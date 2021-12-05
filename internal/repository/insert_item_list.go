package repository

import (
	"blockchain_api/pkg/models"
	"fmt"

	"github.com/sirupsen/logrus"
)

func (p *PostgresDB) InsertItemList(itemList models.BlockchainItem) bool {

	var id int
	query := fmt.Sprintf("INSERT INTO blockchain_list (symbol, price_24h, volume_24h, last_trade_price) values ($1, $2, $3, $4) RETURNING id")

	row := p.DB.QueryRow(query, itemList.Symbol, itemList.Price_24h, itemList.Volume_24h, itemList.Last_trade_price)
	if errSearchItem := row.Scan(&id); errSearchItem != nil {
		logrus.Fatalf("insert error blockchainitem: %s\n", errSearchItem.Error())
	}

	if id == 0 {
		return false
	} else {
		return true
	}
}
