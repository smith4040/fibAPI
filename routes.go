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
	fib "github.com/t-pwk/go-fibonacci"
)

// setupRouter creates the gin router and adds the handlers
func setupRouter(fibNum *uint) *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.GET("/current", serveCurrentNumber(fibNum))
	router.GET("/next", serveNextNumber(fibNum))
	router.GET("/previous", servePreviousNumber(fibNum))

	return router
}

// serveCurrentNumber serves the current number in the Fibonacci sequence. The current number defaults to 0 on API start up.
func serveCurrentNumber(num *uint) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprint(fib.FibonacciBig(*num)))
	}
	return gin.HandlerFunc(fn)
}

// serveNextNumber serves the next larger number in the Fibonacci sequence.
func serveNextNumber(num *uint) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		n := *num
		n++
		c.String(http.StatusOK, fmt.Sprint(fib.FibonacciBig(n)))
		*num = n
	}
	return gin.HandlerFunc(fn)
}

// servePreviousNumber serves the next smaller number in the Fibonacci sequence, stopping at 0.
// If servePreviousNumber is requested while the current number is 0, a 404 error will be presented.
func servePreviousNumber(num *uint) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		n := *num
		if n == 0 {
			c.String(http.StatusNotFound, fmt.Sprint("API does not return negative numbers."))
		} else {
			n--
			c.String(http.StatusOK, fmt.Sprint(fib.FibonacciBig(n)))
		}
		*num = n
	}
	return gin.HandlerFunc(fn)
}

// waitForShutdown gracefully shuts down server
func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal)
	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
