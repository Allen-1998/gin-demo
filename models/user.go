package models

type LoginReq struct {
	Name     string `form:"name" binding:"required"`
	Password int    `form:"password" binding:"required"`
}
