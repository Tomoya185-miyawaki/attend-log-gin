/*
env用ヘルパー
*/
package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() {
	env := os.Getenv("ENV")
	if env != "production" {
		env = "development"
	}
	godotenv.Load(".env." + env)
	godotenv.Load()
}
