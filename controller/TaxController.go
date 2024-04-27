package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/rittbpt/assessment-tax/Request"
	"github.com/rittbpt/assessment-tax/Respone"
	"github.com/rittbpt/assessment-tax/service"
	"encoding/csv"
	"log"
	"strconv"
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


func (t *TaxController) CalCSV(c echo.Context) error {
	file, err := c.FormFile("taxes")
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to read CSV file: "+err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to open CSV file: "+err.Error())
	}
	defer src.Close()

	csvReader := csv.NewReader(src)
	records, err := csvReader.ReadAll()
	var taxs []request.TaxRequest
	for i := 1; i < len(records); i++ {
		log.Println(records[i][0] )
		totalIncome, _ := strconv.ParseFloat(records[i][0], 64)
		wht, _ := strconv.ParseFloat(records[i][1], 64)
		donation, _ := strconv.ParseFloat(records[i][2], 64)
		allowances := []request.Allowance{
			{
				AllowanceType: "donation",
				Amount:        donation,
			},
		}
		taxRequest := request.TaxRequest{
			TotalIncome: totalIncome,
			WHT:         wht,
			Allowances:  allowances,
		}
		taxs = append(taxs, taxRequest)
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to read CSV records: "+err.Error())
	}
	taxData, err := t.TaxService.CalCSV(taxs)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get tax data: "+err.Error())
	}

	result := respone.TaxCSVResponse{Taxes: taxData}

	// return c.JSON(http.StatusOK, taxData)
	return c.JSON(http.StatusOK, result)

}
