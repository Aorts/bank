package repository

import (
	"github.com/jmoiron/sqlx"
)

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) Create(acc Account) (*Account, error) {
	query := "insert into accounts (customer_id, opening_date, account_type, amount, status) values ($1, $2, $3, $4, $5) RETURNING account_id"
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	res, err := tx.Exec(
		query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)
	if err != nil {
		return nil, err
	}
	//fmt.Println(res.LastInsertId())
	/*result, err := r.db.Exec(
		query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)
	if err != nil {
		return nil, err
	}*/
	id, err := res.LastInsertId()
	acc.AccountID = int(id)

	return &acc, nil
}

func (r accountRepositoryDB) GetAll(customerID int) ([]Account, error) {
	query := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where customer_id=$1"
	accounts := []Account{}
	err := r.db.Select(&accounts, query, customerID)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
