package controller

import (
	"todo-backend/helper"
	"fmt"
	"github.com/labstack/echo/v4"
	"time"
)

func GetRoot(c echo.Context) error {
	return helper.MessageSuccess(c, fmt.Sprintf("TODO App By Maxi %d", time.Now().Year()))
}
