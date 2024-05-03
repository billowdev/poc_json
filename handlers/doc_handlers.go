package handlers

import (
	"poc_json/dto"
	"poc_json/models"
	"poc_json/services"
	"poc_json/utils"

	"github.com/gofiber/fiber/v2"
)

type (
	IDocumentHandlerInfs interface {
		HandleTest(c *fiber.Ctx) error
		HandleGetDocumentVersion(c *fiber.Ctx) error
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
	version := models.DocumentVersionModel{
		Version:     1,
		VersionType: "ORIGINAL",
		Value:       p,
	}
	if err := h.documentSrv.CreateDocumentVersion(&version); err != nil {
		return utils.NewResponse[interface{}](c, "E-000", "Test Failed", nil)
	}
	return utils.NewResponse[interface{}](c, "S-000", "Test successfully", dto.SDocumentVersionResponse{
		ID:          version.ID,
		Version:     version.Version,
		VersionType: version.VersionType,
		Value:       version.Value,
	})
}

// HandleGetDocumentVersion implements IDocumentHandlerInfs.
func (h *handlerDeps) HandleGetDocumentVersion(c *fiber.Ctx) error {
	panic("unimplemented")
}
