package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	// Start timer
	start := time.Now()

	// Process request
	err := c.Next()

	// Stop timer
	stop := time.Now()
	latency := stop.Sub(start)

	// Get request details
	method := c.Method()
	path := c.Path()
	status := c.Response().StatusCode()

	// Log the request
	log.Printf("%s %s %d %v", method, path, status, latency)

	return err
}
