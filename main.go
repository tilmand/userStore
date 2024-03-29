package main

import (
	"flag"
	"log"
	"userStore/auth"
	"userStore/config"
	"userStore/http"
	"userStore/store"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("Can't read config file: %v", err)
	}

	sqlStore, err := store.NewSqlStore(conf)
	if err != nil {
		log.Fatalf("main NewSqlStore err: %v", err)
	}

	middleware := auth.NewAuthMiddleware(sqlStore)

	port := flag.Int("port", 8080, "Port number")
	flag.Parse()

	if err := http.NewServer(middleware, sqlStore, port); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)

		return
	}
	http.Wait()
}
