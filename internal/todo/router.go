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
	_ = ctx.Error(fmt.Errorf("An error occured..."))
	ctx.Abort()
}

type TodoRequest struct {
	Field string `json:"field" binding:"required"`
}

func (t TodoRouter) HandleValidationError(ctx *gin.Context) {
	var r TodoRequest
	if err := ctx.ShouldBind(&r); err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
}

func (t TodoRouter) Register(engine *gin.Engine) {
	engine.GET("/hello", t.HandleHelloWorld)
	engine.GET("/error", t.HandleError)
	engine.POST("/validation-error", t.HandleValidationError)
}
