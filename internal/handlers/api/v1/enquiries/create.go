package enquiries

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) Create(ctx *gin.Context) {
	var enquiryData enquiry
	if err := ctx.ShouldBindJSON(&enquiryData); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err := h.enquiriesService.CreateEnquiry(
		enquiryData.Customer,
		enquiryData.Terms_of_shipment,
		enquiryData.Origin_customs,
		enquiryData.Door_pickup,
		enquiryData.Origin_port,
		enquiryData.Pickup_address,
		enquiryData.Cargo_ready_date,
		enquiryData.Destination_customs,
		enquiryData.Door_delivery,
		enquiryData.Destination_port,
		enquiryData.Delivery_address,
		enquiryData.Hs_code,
		enquiryData.Remarks,
		enquiryData.Cargo,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Enquiry created successfully"})
}