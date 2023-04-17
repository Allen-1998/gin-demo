package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.Static("/static","./static")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	index := r.Group("/")
	indexIn(index)
}
