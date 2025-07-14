package middleware

import (
	"net/http"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
)

var ipAccessMap = make(map[string][]time.Time)
var mu sync.Mutex

const (
	maxRequestsPerHour = 100
)

// LimitIPAccess giới hạn số lần truy cập của 1 IP trong 1 giờ
func LimitIPAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		mu.Lock()
		times := ipAccessMap[ip]
		// Lọc các lần truy cập trong 1 giờ gần nhất
		var recent []time.Time
		for _, t := range times {
			if now.Sub(t) < time.Hour {
				recent = append(recent, t)
			}
		}
		recent = append(recent, now)
		ipAccessMap[ip] = recent
		mu.Unlock()

		if len(recent) > maxRequestsPerHour {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests from this IP. Please try again later."})
			return
		}

		c.Next()
	}
} 