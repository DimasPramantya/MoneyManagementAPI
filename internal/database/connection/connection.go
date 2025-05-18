package connection

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

var (
	DBConnections *sql.DB
	err           error
)

func Initiator() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
	)

	DBConnections, err = sql.Open("postgres", dsn)

	// check connection
	err = DBConnections.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}