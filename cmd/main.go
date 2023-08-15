package main

import (
	"julianmindria/backendCoffee/internal/routers"
	"julianmindria/backendCoffee/pkg"
	"log"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	viper.SetConfigName("web.env")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	database, err := pkg.Postgresdb()
	if err != nil {
		log.Fatal(err)
	}
	router := routers.NewRoute(database)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
