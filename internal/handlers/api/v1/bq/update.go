package bq

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(ctx *gin.Context) {
	var bqData bq

	if err := ctx.ShouldBindJSON(&bqData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(bqData, "data")
	_, err := h.bqService.GetBqByID(bqData.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	errr := h.bqService.UpdateBq(
		bqData.Id, 
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

	if errr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	
	

}