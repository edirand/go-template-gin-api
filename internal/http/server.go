package http

import (
	"github.com/edirand/go-template-gin-api/internal/http/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router interface {
	Register(engine *gin.Engine)
}

func NewServer(routers ...Router) *http.Server {
	baseRouter := gin.New()
	baseRouter.Use(middlewares.GinErrorHandler())
	middlewares.ConfigureGinValidation()
	for _, router := range routers {
		router.Register(baseRouter)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: baseRouter,
	}

	return server
}
