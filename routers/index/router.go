package index

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.Static("/static","./static")
	index := r.Group("/")
	indexIn(index)
}
