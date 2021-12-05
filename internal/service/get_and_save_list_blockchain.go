package service

import (
	"blockchain_api/internal/repository"
	"blockchain_api/pkg/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func GetAndSaveBlockchainList(db *repository.PostgresDB) {
	for {

		resp, errResp := http.Get("https://api.blockchain.com/v3/exchange/tickers")
		if errResp != nil {
			logrus.Fatalf("error when requesting a receipt blockchain list: %s \n", errResp.Error())
		}
		// defer resp.Body.Close()

		var blockchainList []models.BlockchainItem
		if errDecode := json.NewDecoder(resp.Body).Decode(&blockchainList); errDecode != nil {
			logrus.Fatalf("decoding errors to the resp.body: %s \n", errDecode.Error())
		}

		errUpdate := db.UpdateItem(blockchainList)
		if errUpdate != nil {
			logrus.Fatalf("update error blockchainitem: %s\n", errUpdate.Error())
		}

		// wait 30 sec
		time.Sleep(30 * time.Second)
	}
}
