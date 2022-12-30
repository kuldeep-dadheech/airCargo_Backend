package bq

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	temp, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bq, errf := h.bqService.GetBqByID(temp)
	if errf != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(bq.Charge, "asdasdasd")
	ctx.JSON(http.StatusOK, gin.H{"bq": bq})
}
