package middlewares

import (
	"github.com/edirand/go-template-gin-api/internal/http/errors"
	"github.com/gin-gonic/gin"
)

func GinErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		for _, err := range context.Errors {
			_, _ = errors.ResolveProblemDetails(context.Writer, context.Request, err)
		}
	}
}
