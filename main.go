// FibonacciAPI is a web based API that steps through the Fibonacci sequence.
// The API exposes 3 endpoints that can be called via HTTP requests:
// current - returns the current number in the sequence, next - returns the next number in the sequence and
// previous - returns the previous number in the sequence.

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	var fibNumber uint = 0

	// Disable log's color
	gin.DisableConsoleColor()

	// Logging to a file.
	f, err := os.Create("gin.log")
	if err != nil {
		log.Fatal("Failed to create log file.")
	}
	gin.DefaultWriter = io.MultiWriter(f)

	router := setupRouter(&fibNumber)

	// Set up the endless server for graceful restarts
	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	server := endless.NewServer(":8080", router)

	log.Println("Starting server")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
