package main

import (
	"go-boilerplate/config"
	"go-boilerplate/internal/server"
	"go-boilerplate/services/kafka"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	kafkaClient, err := kafka.NewKafka(config.Config.KafkaSeeds, config.Config.KafkaTopics)
	if err != nil {
		panic(err)
	}

	s := server.Init(&kafkaClient)
	s.Run()

}
