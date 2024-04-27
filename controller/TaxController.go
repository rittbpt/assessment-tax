package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/rittbpt/assessment-tax/Request"
	"github.com/rittbpt/assessment-tax/service"
)

type TaxController struct {
	TaxService *service.TaxService
}

func NewTaxController(taxService *service.TaxService) *TaxController {
	return &TaxController{
		TaxService: taxService,
	}
}

func (t *TaxController) ChangeDp(c echo.Context) error {
	var requestBody request.PersernalDeduct
	if err := c.Bind(&requestBody); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	result, err := t.TaxService.ChangeDp(requestBody.Amount)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get tax data: "+err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (t *TaxController) Cal(c echo.Context) error {
	var requestBody request.TaxRequest

	if err := c.Bind(&requestBody); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	taxData, err := t.TaxService.Cal(requestBody)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get tax data: "+err.Error())
	}

	return c.JSON(http.StatusOK, taxData)
}
