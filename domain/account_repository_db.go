package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	errs "github.com/kenethrrizzo/banking/error"
	"github.com/kenethrrizzo/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d *AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "insert into accounts (cu_id, ac_opening_date, ac_type, ac_amount, ac_status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.Type, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account -> " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account -> " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.Id = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbclient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbclient}
}