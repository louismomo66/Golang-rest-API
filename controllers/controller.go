package controllers

import (
	"louis/go_projects/models"
	"louis/go_projects/services"
	"net/http"

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
	var user models.User
	
	if err :=	ctx.ShouldBindJSON(&user); err != nil{
	ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"error":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"succes"})

	}
	
	func (uc *UserController) GetUser(ctx *gin.Context) {
		username := ctx.Param("name")
		user,err := uc.UserService.GetUser(&username)
		if err != nil{
			ctx.JSON(http.StatusBadGateway,gin.H{"error":err.Error()})
			return
		}
		ctx.JSON(http.StatusOK,user)
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
		r.GET("/get/:name",uc.GetUser)
		r.GET("/getall",uc.GetAll)
		r.PATCH("/update",uc.UpdateUser)
		r.DELETE("/delete",uc.DeleteUser)
	}