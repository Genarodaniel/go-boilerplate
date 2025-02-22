package server

import (
	"database/sql"
	"go-boilerplate/config"
	"go-boilerplate/internal/app/user"
	userRepository "go-boilerplate/internal/repository/user"
	"go-boilerplate/internal/services/kafka"
	"go-boilerplate/server/handler/healthcheck"
	userHandler "go-boilerplate/server/handler/user"
	"go-boilerplate/server/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func Init(kafkaClient *kafka.KafkaInterface, db *sql.DB) *gin.Engine {

	//update with config env value
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter))
	router.Use(gin.Recovery())
	router.Use(middleware.TimeoutMiddleware(time.Duration(config.Config.ServerTimeout) * time.Second))

	Router(router, kafkaClient, db)

	return router
}

func Router(e *gin.Engine, kafkaClient *kafka.KafkaInterface, db *sql.DB) {
	v1 := e.Group("/v1")
	userGroup := v1.Group("/user")
	healthCheckGroup := v1.Group("/healthcheck")

	userRepository := userRepository.NewUserRepository(db)
	userService := user.NewUserService(*kafkaClient, userRepository)

	healthcheck.Router(healthCheckGroup)
	userHandler.Router(userGroup, userService)
}
