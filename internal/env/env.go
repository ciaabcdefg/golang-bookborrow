package env

import (
	"os"

	"github.com/joho/godotenv"
)

var Hostname string
var Port string
var DBConnection string

func Init() {
	godotenv.Load()

	Hostname = env("HOSTNAME", "localhost")
	Port = env("PORT", "3000")
	DBConnection = env("DB_CONNECTION", "postgres://postgres:1234@localhost:34000/bookborrow?sslmode=disable")
}

func LoadAndGet(name string, placeholder string) string {
	godotenv.Load()
	return env(name, placeholder)
}

func env(name string, placeholder string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		return placeholder
	}
	return value
}
