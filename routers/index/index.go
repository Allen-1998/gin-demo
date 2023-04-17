package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexIn(r *gin.RouterGroup) {

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "hello",
		})
	})
}
