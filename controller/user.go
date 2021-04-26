package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"golang-echo-kafka-example/config"
	"golang-echo-kafka-example/model"
	"golang-echo-kafka-example/producer"
	"golang-echo-kafka-example/util"
	"net/http"
)

type UserController struct {
	producer producer.Producer
}

func NewUserController(producer producer.Producer) *UserController {
	return &UserController{producer: producer}
}

// SaveUser godoc
// @Summary Create a user
// @Description Create a new user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param user body model.UserInput true "New User"
// @Success 200 {object} model.User
// @Failure 500 {object} handler.APIError
// @Router /signup [post]
func (userController *UserController) SaveUser(c echo.Context) error {
	payload := new(model.UserInput)
	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	user := &model.User{UserInput: payload}
	user.ID = uuid.NewV4().String()

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	userController.producer.Produce(config.UserNotificationTopic, string(data))

	return util.Negotiate(c, http.StatusOK, user)
}
