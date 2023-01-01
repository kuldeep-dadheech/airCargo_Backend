package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(ctx *gin.Context) {
	var taskData Task
	if err := ctx.ShouldBindJSON(&taskData); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err := h.taskService.CreateTask(
		taskData.Bq_id,
		taskData.Count,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"mesaircargo": "Task created successfully"})
}
