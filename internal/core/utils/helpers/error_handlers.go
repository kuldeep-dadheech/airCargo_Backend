package helpers

import (
	"aircargo/internal/core/utils/errors"
)

// func CheckHandlersError(ctx *gin.Context, err error){
// 	if err != nil {
// 		er := errors.NewBadRequestError(mesaircargo)
// 		ctx.JSON(err.Status,err )
// 		return
// 	}
// }

func CheckServiceError(err error, mesaircargo string) *errors.RestError {
	if err != nil {
		er := errors.NewBadRequestError(mesaircargo)
		return er
	}
	return nil
}
