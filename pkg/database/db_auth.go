package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mkdr4/hex/internal/domain"
	"github.com/spf13/viper"
)

type Key domain.AuthKey

func GetAuthkey(db *sqlx.DB, level int) ([]*Key, error) {
	tableName := viper.GetString("database.table_name.api_keys")

	keys := []*Key{}

	if err := db.Select(&keys, fmt.Sprintf("SELECT * FROM %s WHERE level >= %d", tableName, level)); err != nil {
		return keys, err
	}

	return keys, nil
}
