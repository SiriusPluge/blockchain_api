package repository

import (
	"blockchain_api/pkg/models"
	"fmt"
)

func (p *PostgresDB) UpdateItem(itemList []models.BlockchainItem) error {

	lenItemList := len(itemList)

	for i := 0; i < lenItemList; i++ {

		var item models.BlockchainItem
		item.Symbol = itemList[i].Symbol
		item.Price_24h = itemList[i].Price_24h
		item.Volume_24h = itemList[i].Volume_24h
		item.Last_trade_price = itemList[i].Last_trade_price

		query := fmt.Sprintf("UPDATE blockchain_list SET price_24h = $1, volume_24h = $2, last_trade_price = $3 WHERE symbol = $4")

		_, errUpdate := p.DB.Exec(query, item.Price_24h, item.Volume_24h, item.Last_trade_price, item.Symbol)
		if errUpdate != nil {
			return errUpdate
		}
	}

	return nil
}
