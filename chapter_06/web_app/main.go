package main

import (
	"log"
	"web_app/config"
	"web_app/server"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting Runners App")
	log.Println("Initializing configuration")
	config := config.InitConfig("runners")
	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	log.Println("Initializing HTTP Server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}