package main

import (
	"context"
	"fmt"
	"github.com/edirand/go-template-gin-api/internal/http"
	"github.com/edirand/go-template-gin-api/internal/todo"
	"log"
	gohttp "net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := http.NewServer(
		todo.NewTodoRouter(),
	)

	done := make(chan bool, 1)
	go gracefulShutdown(server, done)

	err := server.ListenAndServe()
	if err != nil && err != gohttp.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Println("Graceful shutdown complete.")
}

func gracefulShutdown(apiServer *gohttp.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}
