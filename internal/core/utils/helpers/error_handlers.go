package helpers

import (
	"sagebackend/internal/core/utils/errors"

)

// func CheckHandlersError(ctx *gin.Context, err error){
// 	if err != nil {
// 		er := errors.NewBadRequestError(message)
// 		ctx.JSON(err.Status,err )
// 		return 
// 	}
// }

func CheckServiceError(err error, message string) *errors.RestError{
	if err != nil {
		er := errors.NewBadRequestError(message)
		return er
	}
	return nil
}