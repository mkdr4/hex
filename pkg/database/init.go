package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Init() (*sqlx.DB, error) {
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	tableName := viper.GetString("database.database_name")

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, tableName))

	return db, err
}

func Close(db *sqlx.DB) {
	db.Close()
}
