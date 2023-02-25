package log

import (
	"log"

	"github.com/spf13/viper"
)

func Debug(r, m string) {
	if viper.GetBool("debug") {
		log.Println("[DEBUG]", r, m)
	}
}

func Info(r, m string) {
	log.Println("[INFO ]", r, m)
}

func Err(r, m string) {
	log.Println("[ERROR]", r, m)
}

func Fatal(r, m string) {
	log.Fatal("[FATAL]", r, m)
}
