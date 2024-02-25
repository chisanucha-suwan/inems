package handler

import (
	"inmes/package/models"
	"inmes/package/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       users,
	})
}
