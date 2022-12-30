package users

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler)Hey(ctx *gin.Context){
	ctx.String(http.StatusOK, "Hey bitches!")
}