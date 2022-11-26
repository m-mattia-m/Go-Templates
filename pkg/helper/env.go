package helper

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) (string, error) {
	env, err := os.LookupEnv(key)
	if !err || env == "" {
		return env, errors.New("environment variable '" + key + "' empty or not found")
	}
	return env, nil
}

func init() {
	log.Println("[INIT MAIN]: Status -> Main-init is loaded")
	godotenv.Load()
}
