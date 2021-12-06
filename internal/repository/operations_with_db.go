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
