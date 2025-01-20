package controller

import (
	"leaky-bucket/api/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MessageController struct{}

func NewController() *MessageController {
	return new(MessageController)
}

func (uc MessageController) GetMessages(c echo.Context) error {
	page, size := setPageParams(c)
	messages, err := service.GetAll(page, size)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, messages)
}
func setPageParams(c echo.Context) (int, int) {
	pageQuery := c.QueryParam("page")
	sizeQuery := c.QueryParam("size")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(sizeQuery)
	if err != nil {
		size = 10
	}
	return page, size
}
