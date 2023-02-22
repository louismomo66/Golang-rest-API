package main

import (
	"context"
	"log"
	"fmt"
	"louis/go_projects/controllers"
	"louis/go_projects/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
	userservice services.UserService
	UserController controllers.UserController
	ctx context.Context
	usercollection *mongo.Collection
	mongoclient *mongo.Client
	err error

)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient,err = mongo.Connect(ctx,mongoconn)
	if err != nil{
		log.Fatal(err)
	}
  err = mongoclient.Ping(ctx,readpref.Primary())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("mongo connection established")
	usercollection = mongoclient.Database("userdb").Collection("users")
	userservice = services.NewUserService(usercollection,ctx)
	UserController = controllers.New(userservice)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("v1")
	UserController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}