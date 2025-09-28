package controller

import (
	"net/http"                  

	"github.com/labstack/echo/v4"
	"suretigai/entity"
	"suretigai/usecase"
)

type UserController struct {
	userUsecase *usecase.UserUsecase
}

func NewUserController(u *usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: u,
	}
}

func (c *UserController) RegisterUser(ctx echo.Context) error {
	name := ctx.FormValue("name")
	avatar := ctx.FormValue("avatar")

	user := &entity.User{
		Name: name,
		AvaterURL: avatar,
	}

	createdUser, err := c.userUsecase.HandleRegisterUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, createdUser)
}