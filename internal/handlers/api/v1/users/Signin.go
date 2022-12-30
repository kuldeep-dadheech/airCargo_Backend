package users

import (
	"fmt"
	"net/http"
	"sagebackend/internal/core/utils/errors"


	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)



type signinResponse struct {
	UserId   int    `json:"user_id"`
	UserRole string `json:"role"`
}

type SigninRequest struct {
	Email string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

func (h *Handler) Signin(ctx *gin.Context) {
	fmt.Printf("here!!")
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("bad_request_error!!")
		ctx.JSON(err.Status, err)
		return
	}
	person, exists, err := h.usersService.Fetch(user.Email)

	fmt.Println(person, "perereerr 1")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	if exists == false {
		ctx.JSON(http.StatusInternalServerError, gin.H{"account_doesn't_exists": err.Error})
		return
	}

	er := bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(user.Password))
	if er != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": (fmt.Errorf("wrong password").Error())})
		return
	}

	tokens := &signinResponse{
		person.Id,
		"admin",
	}
	ctx.JSON(http.StatusOK, tokens)
}
