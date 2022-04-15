package handler

import (
	"Desktop/shopi/assignment/model"
	"Desktop/shopi/assignment/service"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Add(ctx *fiber.Ctx) error
	Filter(ctx *fiber.Ctx) error
}

type handler struct {
	service service.Service
}
type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func NewHandler(service service.Service) Handler {
	return handler{service: service}
}

func (h handler) Filter(c *fiber.Ctx) error {
	filter := model.OrderFilterModel{}
	err := c.BodyParser(&filter)

	result, err := h.service.Filter(filter)

	if err != nil {
		return c.Status(400).JSON("Filter internal Service  error")
	}

	return c.Status(200).JSON(Response{Data: result})
}

func (h handler) Add(c *fiber.Ctx) error {
	item := model.Order{}
	err := c.BodyParser(&item)

	if err != nil {
		return c.Status(400).JSON("Body is not Parsed Successfully")
	}

	err = h.service.Add(item)
	if err != nil {
		return c.Status(400).JSON("Add Internal Service error")
	}
	return c.Status(200).JSON("Success")
}
