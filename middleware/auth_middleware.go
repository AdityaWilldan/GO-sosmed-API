package middleware

import (
	"GoSosmed/errorhandler"
	"GoSosmed/helper"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			errorhandler.HandleError(ctx, &errorhandler.UnathorizedError{
				Message: "Unauthorized",
			})
			ctx.Abort()
			return
		}
		userID, err := helper.ValidateToken(tokenString)
		if err != nil {
			errorhandler.HandleError(ctx, &errorhandler.UnathorizedError{
				Message: err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Set("userID", *userID)
		ctx.Next()
	}

}
