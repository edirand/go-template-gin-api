package todo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoRouter struct {
}

func NewTodoRouter() *TodoRouter {
	return &TodoRouter{}
}

func (t TodoRouter) HandleHelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello world !"})
}

func (t TodoRouter) HandleError(ctx *gin.Context) {
	_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("An error occured..."))
}

func (t TodoRouter) Register(engine *gin.Engine) {
	engine.GET("/hello", t.HandleHelloWorld)
	engine.GET("/error", t.HandleError)
}
