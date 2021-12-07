package handlers

import (
	"blockchain_api/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetItemBlockchainHandler(ctx *gin.Context) {

	var input models.BlockchainItemPostgres

	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	item, err := h.services.Blockchain.GetBlockchainItem(input.Symbol)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var res models.Response
	res.Price = item.Price_24h
	res.Volume = item.Volume_24h
	res.Last_trade = item.Last_trade_price

	ctx.JSON(http.StatusOK, map[string]interface{}{
		item.Symbol: res,
	})

}
