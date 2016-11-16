// web.go
package main

import "github.com/gin-gonic/gin"
import "net/http"
import "os"

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("view/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Share adult contents you love!",
		})
	})
	if os.Getenv("PORT") == "" {
		router.Run(":" + "8080")
	} else {
		router.Run(":" + os.Getenv("PORT"))
	}
}
