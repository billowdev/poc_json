package handlers

import (
	"poc_json/models"
	"poc_json/services"
	"poc_json/utils"

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
	p, err := utils.NewJSONB(models.STDocumentVersion{
		PortOfLoading:     "HCM",
		PortOfDestination: "HAN",
		CompanyName:       "TEST",
		Address:           "test",
	})
	if err != nil {
		return err
	}
	if err := h.documentSrv.CreateDocumentVersion(models.DocumentVersionModel{
		Version:     1,
		VersionType: "ORIGINAL",
		Value:       p,
	}); err != nil {
		return utils.NewResponse[interface{}](c, "E-400", "Test successfully", nil)
	}
	return c.JSON(fiber.Map{
		"message": h.documentSrv.GetTest(),
	})
}
