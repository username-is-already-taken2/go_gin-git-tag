package main

import (
	"log/slog"
	"net/http"
	"os"
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

// getEnv get key environment variable if exist, otherwise return defaultValue
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func SetupRouter() *gin.Engine {
	if Version == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// Creates a router without any middleware by default
	// r := gin.Default()
	r := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{})
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

	return r
}

func main() {

	// Print Server information to console
	slog.Info("Application", "Version", Version)
	slog.Info("Application", "BuildTime", BuildTime)
	ginPort := getEnv("PORT", "8080")
	slog.Info("Starting GIN on", "PORT", ginPort)

	// Start gin server
	err := SetupRouter().Run(":" + ginPort)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
