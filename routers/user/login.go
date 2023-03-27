package user

import (
	"net/http"

	"gin-demo/models"

	"github.com/gin-gonic/gin"
)

func loginIn(r *gin.RouterGroup) {

	r.POST("/login", func(c *gin.Context) {
		var loginReq models.LoginReq
		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name":     loginReq.Name,
			"password": loginReq.Password,
			"message":  "success",
		})
	})
}
