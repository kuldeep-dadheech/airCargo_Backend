package bq

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(ctx *gin.Context) {
	var bqData bq
	if err := ctx.ShouldBindJSON(&bqData); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err := h.bqService.CreateBq(
		bqData.Is_quote,
		bqData.Enquiry_Id,
		bqData.Valid_till,
		bqData.Vessel_liner,
		bqData.Vessel_total_transit,
		bqData.Vessel_free_days,
		bqData.Freight_charges_vendor,
		bqData.Freight_other_charges_vendor,
		bqData.Total_buy_amount,
		bqData.Total_sell_amount,
		bqData.Profit_percentage,
		bqData.Charge,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"mesaircargo": "BQ created successfully"})
}
