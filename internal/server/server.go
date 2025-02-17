package server

import (
	"database/sql"
	"go-boilerplate/internal/app/healthcheck"
	"go-boilerplate/internal/app/user"
	userRepository "go-boilerplate/internal/repository/user"
	"go-boilerplate/services/kafka"

	"github.com/gin-gonic/gin"
)

func Init(kafkaClient *kafka.KafkaInterface, db *sql.DB) *gin.Engine {

	//update with config env value
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter))
	router.Use(gin.Recovery())

	Router(router, kafkaClient, db)

	return router
}

func Router(e *gin.Engine, kafkaClient *kafka.KafkaInterface, db *sql.DB) {
	v1 := e.Group("/v1")
	userGroup := v1.Group("/user")
	healthCheckGroup := v1.Group("/healthcheck")

	userRepository := userRepository.NewOrderRepository(db)
	userService := user.NewUserService(*kafkaClient, userRepository)

	healthcheck.Router(healthCheckGroup)
	user.Router(userGroup, userService)
}
