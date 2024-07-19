package repository

import (
	"database/sql"
	"fmt"
	"go-learning/model"
)

type AccountRepository struct {
	connection *sql.DB
}

func NewAccountRepository(connection *sql.DB) AccountRepository {
	return AccountRepository{
		connection: connection,
	}
}

func (pr *AccountRepository) GetAccounts() ([]model.Account, error) {
	query := "SELECT id, taxId, amount FROM accounts"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Account{}, err
	}

	var accountList []model.Account
	var accountObj model.Account

	for rows.Next() {
		err = rows.Scan(
			&accountObj.ID,
			&accountObj.TaxId,
			&accountObj.Amount,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Account{}, err
		}

		accountList = append(accountList, accountObj)
	}

	rows.Close()

	return accountList, nil
}
