package pkg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Postgresdb() (*sqlx.DB, error) {
	host := viper.GetString("database.host")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")
	port := viper.GetString("database.dbport")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbName, port)

	return sqlx.Connect("postgres", config)
}
