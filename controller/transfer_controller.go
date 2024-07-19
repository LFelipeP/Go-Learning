package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-learning/service"
)

type transferController struct {
	transferService service.TransferService
}

func NewTransferController(service service.TransferService) transferController {
	return transferController{
		transferService: service,
	}
}

func (p *transferController) GetTransfers(ctx *gin.Context) {
	transfers, err := p.transferService.GetTransfers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, transfers)
}
