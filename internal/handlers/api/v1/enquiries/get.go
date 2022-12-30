package enquiries

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	temp ,err :=strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	enquiry, errf := h.enquiriesService.GetEnquiryById(temp)
	if errf != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"enquiry": enquiry})
}