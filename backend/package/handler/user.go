package handler

import (
	"inems/package/models"
	"inems/package/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type handler struct {
	Service services.UserService
}

func NewUserHandler(db *sqlx.DB) *handler {
	return &handler{
		Service: services.NewUserService(db),
	}
}

func (h *handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.Service.GetUsers()

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
