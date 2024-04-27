package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/rittbpt/assessment-tax/Respone"
	"github.com/rittbpt/assessment-tax/Request"
	"github.com/rittbpt/assessment-tax/service"
	"log"	
	"reflect"
)

type TaxController struct {
	TaxService *service.TaxService
}

func NewTaxController(taxService *service.TaxService) *TaxController {
	return &TaxController{
		TaxService: taxService,
	}
}

func (t *TaxController) Cal(c echo.Context) error {
	var requestBody request.TaxRequest

	if err := c.Bind(&requestBody); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	taxData, err := t.TaxService.Cal(requestBody)
	log.Println(reflect.TypeOf(taxData))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get tax data: "+err.Error())
	}

	response := respone.TaxResponse{
		Tax: float64(taxData),
	}

	log.Println(reflect.TypeOf(response.Tax) , response.Tax)
	return c.JSON(http.StatusOK, response)
}
