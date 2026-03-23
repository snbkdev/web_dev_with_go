package server

import (
	"database/sql"
	"log"

	"github.com/spf13/viper"
)

func InitDatabase(config *viper.Viper) *sql.DB {
	connectionString := config.GetString("database.connection_string")
	maxIdleConnections := config.GetInt("database.max_idle_connections")
	maxOpenConnections := config.GetInt("database.max_open_connections")
	connectionMaxLiftime := config.GetDuration("database.connection_max_lifetime")
	driverName := config.GetString("database.driver_name")
	if connectionString == ""{
		log.Fatalf("Database connection string is missing")
	}
	dbHanler, err := sql.Open(driverName, connectionString)
	if err != nil {
		log.Fatalf("Error while initializing database: %v", err)
	}

	dbHanler.SetMaxIdleConns(maxIdleConnections)
	dbHanler.SetMaxOpenConns(maxOpenConnections)
	dbHanler.SetConnMaxLifetime(connectionMaxLiftime)
	err = dbHanler.Ping()
	if err != nil {
		dbHanler.Close()
		log.Fatalf("Error while validating database: %v", err)
	}

	return dbHanler
}