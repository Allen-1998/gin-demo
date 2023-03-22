package user

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	user := r.Group("/user")
	loginIn(user)
}
