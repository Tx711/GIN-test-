package main



import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/catalog", func(c *gin.Context) {
		c.String(http.StatusOK, "CATALOG <3 \ntoo much wow")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}