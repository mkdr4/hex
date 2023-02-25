package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mkdr4/hex/internal/domain"
	"github.com/spf13/viper"
)

type (
	Cryptocurrencies domain.Cryptocurrencies
	Wallets          domain.Wallet
	Networks         domain.Networks
)

const (
	crpts = `name, symbol, image_link, course`
	wlts  = `wallets`
	ntwks = `networks`
)

func GetCryptocurrencies(db *sqlx.DB) ([]*Cryptocurrencies, error) {
	tableName := viper.GetString("database.table_name.crypto_data")

	crypto := []*Cryptocurrencies{}

	if err := db.Select(&crypto, fmt.Sprintf("SELECT %s FROM %s", crpts, tableName)); err != nil {
		return crypto, err
	}

	return crypto, nil
}

func GetWallets(db *sqlx.DB, cryptoName string) (*[]byte, error) {
	tableName := viper.GetString("database.table_name.crypto_data")

	wallets := &[]byte{}

	if err := db.Get(wallets, fmt.Sprintf("SELECT %s FROM %s WHERE name = '%s' ", wlts, tableName, cryptoName)); err != nil {
		return nil, err
	} else if len(*wallets) == 0 {
		return nil, fmt.Errorf("not found wallets for %s", cryptoName)
	}

	return wallets, nil
}

func GetNetworks(db *sqlx.DB, cryptoName string) (*[]byte, error) {
	tableName := viper.GetString("database.table_name.crypto_data")

	networks := &[]byte{}

	if err := db.Get(networks, fmt.Sprintf("SELECT %s FROM %s WHERE name = '%s' ", ntwks, tableName, cryptoName)); err != nil {
		return nil, err
	} else if len(*networks) == 0 {
		return nil, fmt.Errorf("not found networks for %s", cryptoName)
	}

	return networks, nil
}
