package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-learning/service"
)

type accountController struct {
	accountService service.AccountService
}

func NewAccountController(service service.AccountService) accountController {
	return accountController{
		accountService: service,
	}
}

func (p *accountController) GetAccounts(ctx *gin.Context) {
	accounts, err := p.accountService.GetAccounts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, accounts)
}
