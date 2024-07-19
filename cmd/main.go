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
	TransferRepository := repository.NewTransferRepository(dbConnection)

	//Services
	AccountService := service.NewAccountService(AccountRepository)
	TransferService := service.NewTransferService(TransferRepository)

	//Contorllers
	AccountController := controller.NewAccountController(AccountService)
	TransferController := controller.NewTransferController(TransferService)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/accounts", AccountController.GetAccounts)
	server.GET("/transfers", TransferController.GetTransfers)

	server.Run(":8000")
}
