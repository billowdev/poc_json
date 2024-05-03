package main

import "github.com/gofiber/fiber/v2"

type (
	IHandlerInfs interface {
		HandleTest(c *fiber.Ctx)
	}
	handlerDeps struct {
		services IServicesInfs
	}
)

func NewHandler(services IServicesInfs) IHandlerInfs {
	return &handlerDeps{
		services: services,
	}
}

func (h *handlerDeps) HandleTest(c *fiber.Ctx) {
	c.JSON(fiber.Map{
		"message": h.services.GetTest(),
	})
}
