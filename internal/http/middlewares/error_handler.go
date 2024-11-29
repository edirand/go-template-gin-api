package middlewares

import (
	"github.com/edirand/go-template-gin-api/internal/http/errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strings"
)

func GinErrorHandler() gin.HandlerFunc {
	registerCustomErrorsMappers()

	return func(context *gin.Context) {
		context.Next()
		for _, err := range context.Errors {
			_, _ = errors.ResolveProblemDetails(context.Writer, context.Request, err)
		}
	}
}

func registerCustomErrorsMappers() {
	errors.Map[validator.ValidationErrors](func() errors.ProblemDetailErr {
		return &errors.ProblemDetail{
			Status: http.StatusBadRequest,
			Detail: "One or more validation errors occurred.",
		}
	})
}

func ConfigureGinValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
