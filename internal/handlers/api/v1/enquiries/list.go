package enquiries

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) List(ctx *gin.Context) {
	enquiries, err := h.enquiriesService.GetAllEnquiries()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"enquiries": enquiries})
}