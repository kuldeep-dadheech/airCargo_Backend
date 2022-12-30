package bq

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) List(ctx *gin.Context) {
	bqs, err := h.bqService.GetAllBq()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(bqs, "bqqqqqllll")
	ctx.JSON(http.StatusOK, gin.H{"bqs": bqs})
}
type ListResponse struct {
	
}