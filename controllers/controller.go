package controllers

import (
	"louis/go_projects/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController{
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	 ctx.JSON(200,"")
	}
	
	func (uc *UserController) GetUser(ctx *gin.Context) {
		ctx.JSON(200,"")
	}
	
	func (uc *UserController) GetAll(ctx *gin.Context) {
		ctx.JSON(200,"")
	}
	
	func (uc *UserController) UpdateUser(ctx *gin.Context)  {
		ctx.JSON(200,"")
	}
	
	func (uc *UserController) DeleteUser(ctx *gin.Context) {
		ctx.JSON(200,"")
	}

	func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup){
		r := rg.Group("/user")
		r.POST("/create",uc.CreateUser)
	}