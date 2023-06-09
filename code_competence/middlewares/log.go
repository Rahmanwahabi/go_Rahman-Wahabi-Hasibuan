package middlewares

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}

	logger := log.New(logFile, "", log.LstdFlags)

	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		logger.Printf("[%s] %s %s %s %d %v", end.Format("2006/01/02 - 15:04:05"), clientIP, method, path, statusCode, latency)
	}
}
