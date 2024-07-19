package main

import (
	"go-learning/controller"
	"go-learning/db"
	"go-learning/repository"
	"go-learning/service"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//Db
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	//Repos
	AccountRepository := repository.NewAccountRepository(dbConnection)

	//Services
	AccountService := service.NewAccountService(AccountRepository)

	//Contorllers
	AccountController := controller.NewAccountController(AccountService)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/accounts", AccountController.GetAccounts)

	server.Run(":8000")
}
