package main

import (
	"go-boilerplate/config"
	"go-boilerplate/database"
	"go-boilerplate/internal/server"
	"go-boilerplate/services/kafka"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	db := database.Connect()
	database.Migrate(db)

	defer db.Close()

	kafkaClient, err := kafka.NewKafka(config.Config.KafkaSeeds)
	if err != nil {
		panic(err)
	}

	s := server.Init(&kafkaClient, db)
	s.Run()

}
