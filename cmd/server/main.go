package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"stkpush-go/internal/handler"
)

const defaultPort = "4000"

func main() {
	// Determine server port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Use release mode in production
	gin.SetMode(gin.ReleaseMode)

	// Initialize router
	router := gin.New()
	router.Use(gin.Recovery()) // panic recovery middleware
	router.GET("/stk", handler.STKPushHandler)

	// HTTP server configuration
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	// Run server in a separate goroutine
	go func() {
		log.Printf("Server started on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server encountered an error: %v", err)
		}
	}()

	// Wait for termination signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown signal received, shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
