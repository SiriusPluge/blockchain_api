package repository

import (
	"blockchain_api/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// Backend инкапсулирует подключение к MongoDB
type PostgresDB struct {
	DB *sqlx.DB
}

type ConfigPostgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// ConnectPostgresDB is used to connect the Postgres Database
func NewConnectionPostgresDB(cfg *ConfigPostgres) *PostgresDB {

	fmt.Printf("initializing the connection to the database in: %s \n", cfg.Port)

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		logrus.Fatalf("error open postgresDB: %s", err.Error())
	}
	// defer db.Close()

	errConnDB := db.Ping()
	if errConnDB != nil {
		logrus.Fatalf("error connection postgresDB: %s", errConnDB.Error())
	}

	fmt.Printf("successful connection to the database in: %s", cfg.Port)

	// insert to the DB blockchainList
	initInsertBCList(db)

	return &PostgresDB{DB: db}
}

func initInsertBCList(db *sqlx.DB) {

	// get getting data from https://api.blockchain.com/v3/exchange/tickers
	resp, errResp := http.Get("https://api.blockchain.com/v3/exchange/tickers")
	if errResp != nil {
		logrus.Fatalf("error when requesting a receipt blockchain list: %s \n", errResp.Error())
	}
	defer resp.Body.Close()

	//decoding of the received data
	var blockchainList []models.BlockchainItem
	if errDecode := json.NewDecoder(resp.Body).Decode(&blockchainList); errDecode != nil {
		logrus.Fatalf("decoding errors to the resp.body: %s \n", errDecode.Error())
	}

	// writing data to the database
	_, errExec := db.NamedExec(`INSERT INTO blockchain_list (symbol, price_24h, volume_24h, last_trade_price)
	    VALUES (:symbol, :price_24h, :volume_24h, :last_trade_price)`, blockchainList)
	if errExec != nil {
		logrus.Fatalf("insert error blockchainitem: %s\n", errExec.Error())
	}
}
