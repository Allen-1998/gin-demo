package user

import (
	"net/http"

	"gin-demo/models"

	"github.com/gin-gonic/gin"
)

func loginIn(r *gin.RouterGroup) {

	r.POST("/login", func(c *gin.Context) {
		var loginReq models.LoginReq
		err := c.Bind(&loginReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求参数异常",
			})
			return
		}
		name := c.PostForm("name")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"name":     name,
			"password": password,
			"message":  "success",
		})
	})
}
