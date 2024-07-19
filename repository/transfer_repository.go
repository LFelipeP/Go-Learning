package repository

import (
	"database/sql"
	"fmt"
	"go-learning/model"
)

type TransferRepository struct {
	connection *sql.DB
}

func NewTransferRepository(connection *sql.DB) TransferRepository {
	return TransferRepository{
		connection: connection,
	}
}

func (pr *TransferRepository) GetTransfers() ([]model.Transfer, error) {
	query := "SELECT id, originAccountId, destinationAccountId, amount FROM transfers"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Transfer{}, err
	}

	var transferList []model.Transfer
	var transferObj model.Transfer

	for rows.Next() {
		err = rows.Scan(
			&transferObj.ID,
			&transferObj.OriginAccountId,
			&transferObj.DestinationAccountId,
			&transferObj.Amount,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Transfer{}, err
		}

		transferList = append(transferList, transferObj)
	}

	rows.Close()

	return transferList, nil
}
