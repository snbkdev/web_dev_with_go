package main

import (
	"log"
	"os"
	"web_app/config"
	"web_app/server"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting Runners App")
	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())
	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	log.Println("Initializing Prometheus")
	go server.InitPrometheus()
	log.Println("Initializing HTTP Server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}

func getConfigFileName() string {
	env := os.Getenv("ENV")
	if env != "" {
		return "runners-" + env
	}

	return "runners"
}