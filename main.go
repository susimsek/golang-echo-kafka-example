package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang-echo-kafka-example/config"
	"golang-echo-kafka-example/consumer"
	"golang-echo-kafka-example/controller"
	_ "golang-echo-kafka-example/docs"
	"golang-echo-kafka-example/handler"
	"golang-echo-kafka-example/producer"
	"golang-echo-kafka-example/routes"
	"golang-echo-kafka-example/util"
)

var userController *controller.UserController

// @title Golang User REST API
// @description Provides access to the core features of Golang User REST API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func main() {
	e := echo.New()

	e.HTTPErrorHandler = handler.ErrorHandler
	e.Validator = util.NewValidationUtil()
	config.CORSConfig(e)

	routes.GetUserApiRoutes(e, userController)
	routes.GetSwaggerRoutes(e)

	// echo server 9000 de başlatıldı.
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ServerPort)))
}

func init() {
	p := config.InitKafkaProducer()
	producer := producer.NewProducer(p)
	userController = controller.NewUserController(producer)
	c := config.InitKafkaConsumer(config.UserConsumerGroup)
	consumer := consumer.NewConsumer(c)
	go consumer.Consume([]string{config.UserNotificationTopic})
}
