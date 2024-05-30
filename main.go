package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// VersionInfo holds version information
type VersionInfo struct {
	Version   string `json:"version"`
	BuildTime string `json:"build_time"`
	Timestamp string `json:"timestamp"`
}

// versionMiddleware is a Gin middleware that attaches versionInfo to the context
func versionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create the version info struct
		versionInfo := VersionInfo{
			Version:   Version,
			BuildTime: BuildTime,
			Timestamp: time.Now().Format(time.RFC3339),
		}
		// Attach versionInfo to the context
		c.Set("versionInfo", versionInfo)
		// Pass control to the next handler
		c.Next()
	}
}

// getVersionInfoHandler returns the version information
func getVersionInfoHandler(c *gin.Context) {
	// Create the version info struct
	versionInfo := VersionInfo{
		Version:   Version,
		BuildTime: BuildTime,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	// Return version info as a JSON response
	c.JSON(http.StatusOK, versionInfo)
}

func main() {
	r := gin.Default()

	// Simple Default Page
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})
	// Simple Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "OK",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// Route to default versionInfo
	r.GET("/version", func(c *gin.Context) {
		c.Request.URL.Path = "/version/handler"
		r.HandleContext(c)
	})

	// Route to get version info from middleware
	r.GET("/version/middleware", versionMiddleware(), func(c *gin.Context) {
		versionInfo, exists := c.Get("versionInfo")
		if exists {
			c.JSON(http.StatusOK, versionInfo)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Version info not found"})
		}
	})

	// Route to get versionInfo from a handler
	r.GET("/version/handler", getVersionInfoHandler)

	// Start gin server
	r.Run()
}
