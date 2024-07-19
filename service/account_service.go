package service

import (
	"go-learning/model"
	"go-learning/repository"
)

type AccountService struct {
	repository repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return AccountService{
		repository: repo,
	}
}

func (ps *AccountService) GetAccounts() ([]model.Account, error) {
	return ps.repository.GetAccounts()
}
