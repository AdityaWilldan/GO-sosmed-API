package errorhandler

import (
	"GoSosmed/dto"
	"GoSosmed/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(ctx *gin.Context, err error) {
	var StatusCode int

	switch err.(type) {
	case *NotFoundError:
		StatusCode = http.StatusNotFound
	case *BadRequestError:
		StatusCode = http.StatusBadRequest
	case *InternalServerError:
		StatusCode = http.StatusInternalServerError
	case *UnathorizedError:
		StatusCode = http.StatusUnauthorized
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: StatusCode,
		Message:    err.Error(),
	})

	ctx.JSON(StatusCode, response)
}
