package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(ctx *gin.Context) {
	var id int
	id, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		ctx.JSON(http.StatusBadRequest, errParam)
		return
	}
	person, exists, e := h.usersService.FetchUserById(id)
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": (fmt.Errorf("user does not exists")).Error()})
	}
	if e != nil {
		ctx.JSON(http.StatusInternalServerError, e)
		return 
	}

	ctx.JSON(200, person)
	return
}


