package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(ctx *gin.Context) {
	var taskData Task

	if err := ctx.ShouldBindJSON(&taskData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.taskService.SelectTask(taskData.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	errr := h.taskService.UpdateTask(
		taskData.Id,
		taskData.Bq_id,
		taskData.Count,
	)

	if errr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}