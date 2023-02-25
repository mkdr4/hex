package log

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

const (
	// prefix database
	DB = "db:"
	// prefix server
	S = "server:"
)

func Setup() (*os.File, error) {
	f, err := os.OpenFile(viper.GetString("logfile"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return f, err
	}

	if !viper.GetBool("debug") {
		log.SetOutput(f)
	}

	return f, nil
}
