package handlers

import (
	"poc_json/services"

	"github.com/gofiber/fiber/v2"
)

type (
	IDocumentHandlerInfs interface {
		HandleTest(c *fiber.Ctx) error
	}
	handlerDeps struct {
		documentSrv services.IDocumentSrvInfs
	}
)

func NewDocumentHandler(documentSrv services.IDocumentSrvInfs) IDocumentHandlerInfs {
	return &handlerDeps{
		documentSrv: documentSrv,
	}
}

func (h *handlerDeps) HandleTest(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": h.documentSrv.GetTest(),
	})
}
