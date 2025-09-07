package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"heloo-go/internal/domain"
	"heloo-go/internal/service"
)

type ItemHandler struct {
	service service.ItemService
}

func NewItemHandler(s service.ItemService) *ItemHandler { return &ItemHandler{service: s} }

func (h *ItemHandler) List(c echo.Context) error {
	items, err := h.service.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		id = c.QueryParam("id")
	}
	item, err := h.service.Get(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) Create(c echo.Context) error {
	var in domain.Item
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid json"})
	}
	created, err := h.service.Create(in)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *ItemHandler) Update(c echo.Context) error {
	var in domain.Item
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid json"})
	}
	updated, err := h.service.Update(in)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "not found" {
			status = http.StatusNotFound
		}
		return c.JSON(status, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *ItemHandler) Delete(c echo.Context) error {
	id := c.QueryParam("id")
	if err := h.service.Delete(id); err != nil {
		status := http.StatusBadRequest
		if err.Error() == "not found" {
			status = http.StatusNotFound
		}
		return c.JSON(status, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
