package domain

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	errs "github.com/kenethrrizzo/banking/error"
	"github.com/kenethrrizzo/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) FindById(id string) (*Account, *errs.AppError) {
	var a Account

	sqlSelect := "select Id, CustomerId, OpeningDate, Type, Amount, Status from	Accounts where Id = ?"

	err := d.client.Get(&a, sqlSelect, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(fmt.Sprintf("Account with id %s not found: %s", id, err.Error()))
			return nil, errs.NewNotFoundError(fmt.Sprintf("Account with id %s not found", id))
		}
		logger.Error("Error while scanning account from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Error while scanning account from database")
	}

	logger.Info("Account found")
	logger.Debug(fmt.Sprint("Account:", a))
	return &a, nil
} 

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "insert into Accounts (CustomerId, OpeningDate, Type, Amount, Status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.Type, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	a.Id = strconv.FormatInt(id, 10)
	logger.Info(fmt.Sprintf("Account with id %s saved", a.Id))
	logger.Debug(fmt.Sprint("Account:", a))
	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	result, err := tx.Exec("insert into Transactions (AccountId, Amount, Type, Date) values (?, ?, ?, ?)", t.AccountId, t.Amount, t.Type, t.Date)
	if err != nil {
		logger.Error("Error while executing insert in table Transactions: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	if t.IsWithdrawal() {
		_, err = tx.Exec("update Accounts set Amount = Amount - ? where Id = ?", t.Amount, t.AccountId)
	} else if t.IsDeposit() {
		_, err = tx.Exec("update Accounts set Amount = Amount + ? where Id = ?", t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	account, apperr := d.FindById(t.AccountId)
	if apperr != nil {
		return nil, apperr
	}

	t.Id = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount

	logger.Debug(fmt.Sprint("Transaction: ", t))

	return &t, nil
}

func NewAccountRepositoryDb(dbclient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbclient}
}