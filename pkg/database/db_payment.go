package database

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mkdr4/hex/internal/domain"
	"github.com/spf13/viper"
)

type Payment domain.Payment

const values = `
:id, :store, :signature, :fiat, :fiat_currency, 
:crypto, :crypto_currency, :pay_description, 
:redirect_url, :webhook_url, :expiration_time, :status
`

const update = `
crypto=:crypto
`

/*id=:id, store=:store, signature=:signature, fiat=:fiat,
fiat_currency=:fiat_currency, crypto=:crypto, crypto_currency=:crypto_currency,
pay_description=:pay_description, redirect_url=:redirect_url,
webhook_url=:webhook_url, expiration_time=:expiration_time, status=:status */

func CreatePayment(db *sqlx.DB, data []byte) error {
	tableName := viper.GetString("database.table_name.payments")

	var pay *Payment
	err := json.Unmarshal(data, pay)
	if err != nil {
		return err
	}

	if _, err = db.NamedExec(fmt.Sprintf("INSERT INTO %s VALUES (%s)", tableName, values), pay); err != nil {
		return err
	}

	return nil
}

func UpdatePayment(db *sqlx.DB, data []byte) error {
	tableName := viper.GetString("database.table_name.payments")

	var pay *Payment
	err := json.Unmarshal(data, pay)
	if err != nil {
		return err
	}

	if _, err = db.NamedExec(fmt.Sprintf("UPDATE %s SET %s WHERE id = '%s' ", tableName, update, pay.Id), pay); err != nil {
		return err
	}

	return nil
}

func GetPayment(db *sqlx.DB, id string) ([]*Payment, error) {
	tableName := viper.GetString("database.table_name.payments")
	pay := []*Payment{}

	if err := db.Select(&pay, fmt.Sprintf("SELECT * FROM %s WHERE id = '%s' ", tableName, id)); err != nil {
		return pay, err
	}
	return pay, nil
}

func DeletePayment(db *sqlx.DB, id string) error {
	tableName := viper.GetString("database.table_name.payments")

	if _, err := db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = '%s' ", tableName, id)); err != nil {
		return err
	}

	return nil
}
