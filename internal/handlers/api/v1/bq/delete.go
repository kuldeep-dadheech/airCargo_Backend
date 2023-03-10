package bq

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(ctx *gin.Context) {
	id, errParam := strconv.Atoi(ctx.Param("id"))

	if errParam != nil {
		ctx.JSON(http.StatusBadRequest, errParam)
	}

	err := h.bqService.DeleteBq(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"mesaircargo": "BQ deleted successfully"})
}
