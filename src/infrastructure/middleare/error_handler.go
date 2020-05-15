package middleare

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"movies/src/domain/exception"
	"movies/src/infrastructure/utils/logger"
	"movies/src/infrastructure/utils/rest_errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		err := context.Errors.Last()
		if err == nil {
			return
		}

		// Use reflect.TypeOf(err.Err) to known the type of your error
		if _, ok := errors.Cause(err.Err).(exception.Validations); ok {
			restErr := rest_errors.NewBadRequestError(err.Error())
			logger.Error(restErr.Message(), restErr)
			context.JSON(restErr.Status(), restErr)
			return
		}

		// Use reflect.TypeOf(err.Err) to known the type of your error
		if _, ok := errors.Cause(err.Err).(exception.MovieNotFound); ok {
			restErr := rest_errors.NewNotFoundError(err.Error())
			logger.Error(restErr.Message(), restErr)
			context.JSON(restErr.Status(), restErr)
			return
		}

		restErr := rest_errors.NewInternalServerError(err.Error(), err)
		logger.Error(restErr.Message(), restErr)
		context.JSON(restErr.Status(), restErr)
	}
}
