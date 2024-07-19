package service

import (
	"go-learning/model"
	"go-learning/repository"
)

type TransferService struct {
	repository repository.TransferRepository
}

func NewTransferService(repo repository.TransferRepository) TransferService {
	return TransferService{
		repository: repo,
	}
}

func (ps *TransferService) GetTransfers() ([]model.Transfer, error) {
	return ps.repository.GetTransfers()
}
