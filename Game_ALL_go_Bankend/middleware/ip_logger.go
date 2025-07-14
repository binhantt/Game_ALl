package middleware

import (
	"os"
	"time"
	"github.com/gin-gonic/gin"
)

// LogIPToFile ghi IP truy cập vào file logs/ip_access.log
func LogIPToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now().Format("2006-01-02 15:04:05")
		logLine := now + " - " + ip + "\n"

		os.MkdirAll("logs", 0755)
		f, err := os.OpenFile("logs/ip_access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			f.WriteString(logLine)
			f.Close()
		}
		c.Next()
	}
} 