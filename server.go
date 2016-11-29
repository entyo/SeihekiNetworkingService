// web.go
package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize server
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("view/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Share adult contents you love!",
		})
	})

	// For Let’s Encrypt
	router.GET("/.well-known/acme-challenge/:value", func(c *gin.Context) {
		req := os.Getenv("LETSENCRYPT_REQUEST")
		res := os.Getenv("LETSENCRYPT_RESPONSE")

		if c.Param("value") == req {
			c.String(http.StatusOK, res)
		}
	})

	if os.Getenv("PORT") == "" {
		router.Run(":" + "8080")
	} else {
		router.Run(":" + os.Getenv("PORT"))
	}
}
