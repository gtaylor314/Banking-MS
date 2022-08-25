package domain

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/gtaylor314/Banking-MS/errs"
	"github.com/gtaylor314/Banking-MS/logger"

	"github.com/jmoiron/sqlx"
)

// TransactionRepositoryDb is an adapter which implements the port TransactionRepository
type TransactionRepositoryDb struct {
	db_conn *sqlx.DB
}

func (t TransactionRepositoryDb) SaveTrans(tran Transaction) (*Transaction, *errs.AppError) {
	// if the transaction is successfully inserted but the account table fails to update, we want to be able to rollback
	// all changes to the database - to do this, we use a sql transaction
	tx, err := t.db_conn.DB.Begin()
	if err != nil {
		logger.Error("error while creating database transaction tx" + err.Error())
		return nil, errs.UnexpectedErr("unexpected error creating database transaction tx")
	}

	// here we defer the sql transaction rollback - if the sql transaction completes, it is committed before SaveTrans exits
	// and the deferred rollback call becomes a no-op - if the sql transaction fails, it won't be committed and the rollback
	// is called as SaveTrans exits
	defer tx.Rollback()

	// define the SQL Insert command - transactions is the name of the table in our banking db
	sqlInsertCmd := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"

	// Exec() executes the query/command using the query string and arguments - returns no rows
	result, err := tx.Exec(sqlInsertCmd, tran.AccountID, tran.Amount, tran.TransactionType, tran.TransactionDate)
	if err != nil {
		logger.Error("error while inserting transaction" + err.Error())
		return nil, errs.UnexpectedErr("unexpected database error during transaction insert")
	}
	// LastInsertId() returns the last integer generated by executing a database command (e.g. the transaction id
	// auto-generated) when inserting a transaction into the transactions table
	tranId, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while retrieving the transaction ID" + err.Error())
		return nil, errs.UnexpectedErr("unexpected database error while retrieving the transaction id")
	}

	// FormatInt returns the string representation of an int64 variable (tranId) using the given base (in this case base 10)
	tran.TransactionID = strconv.FormatInt(tranId, 10)

	// UpdateAcctAmount() updates the accounts table based on if the transaction type was a withdrawal or a deposit using
	// the amount of the transaction
	transaction, appErr := t.UpdateAcctAmount(tran, tx)
	if appErr != nil {
		return nil, appErr
	}

	if err = tx.Commit(); err != nil {
		logger.Error("error while committing the transaction to the database" + err.Error())
		return nil, errs.UnexpectedErr("unexpected database error while committing the transaction")
	}

	return transaction, nil
}

func (t TransactionRepositoryDb) GetAcctAmount(acctId string) (*float64, *errs.AppError) {
	// define the SQL query to get the account amount (or balance) via the account ID
	sqlSelAccountAmount := "SELECT amount FROM accounts WHERE account_id = ?"

	var amount float64
	err := t.db_conn.Get(&amount, sqlSelAccountAmount, acctId)
	if err != nil {
		logger.Error("error while retrieving account's balance" + err.Error())
		if err == sql.ErrNoRows {
			// the provided account id wasn't found in the accounts table
			return nil, errs.NotFoundErr("account id provided was not found in database")
		}
		return nil, errs.UnexpectedErr("unexpected database error while retrieving the account's current balance")
	}
	return &amount, nil
}

func (t TransactionRepositoryDb) UpdateAcctAmount(tran Transaction, tx *sql.Tx) (*Transaction, *errs.AppError) {
	// if the transaction type is a withdrawal
	if strings.ToLower(tran.TransactionType) == "withdrawal" {
		// update AcctAmount (account balance) by subtracting the amount withdrawn
		tran.AcctAmount = tran.AcctAmount - tran.Amount
		sqlUpdateAcctAmount := "UPDATE accounts SET amount = ? WHERE account_id = ?"
		_, err := tx.Exec(sqlUpdateAcctAmount, tran.AcctAmount, tran.AccountID)
		if err != nil {
			logger.Error("error while updating account's balance" + err.Error())
			return nil, errs.UnexpectedErr("unexpected database error while updating account's balance")
		}
		return &tran, nil
	}
	// otherwise, the transaction type is a deposit
	// update AcctAmount (account balance) by adding the amount deposited
	tran.AcctAmount = tran.AcctAmount + tran.Amount
	sqlUpdateAcctAmount := "UPDATE accounts SET amount = ? WHERE account_id = ?"
	_, err := tx.Exec(sqlUpdateAcctAmount, tran.AcctAmount, tran.AccountID)
	if err != nil {
		logger.Error("error while updating account's balance" + err.Error())
		return nil, errs.UnexpectedErr("unexpected database error while updating account's balance")
	}
	return &tran, nil
}

// NewTransactionRepositoryDb() takes in a db connection and returns a TransactionRepositoryDb object
func NewTransactionRepositoryDb(db_conn *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{db_conn: db_conn}
}
