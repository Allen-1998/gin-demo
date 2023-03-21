package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginIn(r *gin.Engine) {
	user := r.Group("/user")

	user.POST("/login", func(c *gin.Context) {
		name := c.PostForm("name")
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "name is not defined",
			})
			return
		}
		password := c.PostForm("password")
		if password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "password is not defined",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name":     name,
			"password": password,
			"message":  "success",
		})
	})
}
