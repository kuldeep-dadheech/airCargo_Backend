package enquiries

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(ctx *gin.Context) {
	var enquiryData enquiry

	if err := ctx.ShouldBindJSON(&enquiryData); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	enquiry, err := h.enquiriesService.GetEnquiryById(enquiryData.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("enquiry", enquiry)

	errf := h.enquiriesService.UpdateEnquiry(enquiryData.Id, enquiryData.Customer, enquiryData.Terms_of_shipment, enquiryData.Origin_customs, enquiryData.Door_pickup, enquiryData.Origin_port, enquiryData.Pickup_address, enquiryData.Cargo_ready_date, enquiryData.Destination_customs, enquiryData.Door_delivery, enquiryData.Destination_port, enquiryData.Delivery_address, enquiryData.Hs_code, enquiryData.Remarks, enquiryData.Cargo)

	if errf != nil {
		ctx.JSON(http.StatusInternalServerError, errf)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mesaircargo": "Enquiry updated successfully"})
}
