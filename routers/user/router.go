package user

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	loginIn(r)
}
