package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"stkpush-go/internal/handler"
)

func main() {
	r := gin.Default()
	r.GET("/stk", handler.STKPushHandler)

	log.Println("🚀 Server running on port 4000")
	if err := r.Run(":4000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
