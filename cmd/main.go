package main

import (
	"go-boilerplate/config"
	"go-boilerplate/database"
	"go-boilerplate/internal/server"
	"go-boilerplate/services/kafka"
	"log"
	"os"
	"os/signal"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	database.Migrate(db)

	kafkaClient, err := kafka.NewKafka(config.Config.KafkaSeeds)
	if err != nil {
		panic(err)
	}

	atInterruption(func() {
		log.Printf("closing DB connection")
		db.Close()
		log.Printf("DB connection closed")
		log.Printf("Server shutdown")
	})

	s := server.Init(&kafkaClient, db)
	s.Run()

}

func atInterruption(fn func()) {
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc

		fn()
		os.Exit(1)
	}()
}
