package server

import (
	"go-boilerplate/internal/domain/healthcheck"
	"go-boilerplate/internal/domain/user"
	"go-boilerplate/services/kafka"

	"github.com/gin-gonic/gin"
)

func Init(kafkaClient *kafka.KafkaInterface) *gin.Engine {

	//update with config env value
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter))
	router.Use(gin.Recovery())

	Router(router, kafkaClient)

	return router
}

func Router(e *gin.Engine, kafkaClient *kafka.KafkaInterface) {
	v1 := e.Group("/v1")
	userGroup := v1.Group("/user")
	healthCheckGroup := v1.Group("/healthcheck")

	userService := user.NewuserService(*kafkaClient)

	healthcheck.Router(healthCheckGroup)
	user.Router(userGroup, userService)
}
