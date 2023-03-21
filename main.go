package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 基本路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	// 路由分组
	user := r.Group("/user")
	{
		user.GET("/", func(c *gin.Context) {
			name := c.Query("name")
			c.JSON(http.StatusOK, gin.H{
				"name":    name,
				"message": "success",
			})
		})
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
	return r
}

func main() {
	r := setupRouter()
	// 监听端口，默认在8080
	r.Run()
}
