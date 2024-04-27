package service

import (
	"math"
	_ "log"
	"github.com/rittbpt/assessment-tax/repository"
	"github.com/rittbpt/assessment-tax/Request"
	"github.com/rittbpt/assessment-tax/Respone"
)

type TaxService struct {
	TaxRepo *repository.TaxRepository
}

func NewTaxService(repo *repository.TaxRepository) *TaxService {
	return &TaxService{
		TaxRepo: repo,
	}
}

func calculateTax(totalIncome float64, personalDeduct float64, wht float64, allowances []request.Allowance, kDeduct float64) respone.TaxResponse {
	taxTable := []struct {
		min   float64
		max   float64
		per   float64
		level string
	}{
		{
			min:   0.0,
			max:   150000.0,
			per:   0,
			level: "0-150,000",
		},
		{
			min:   150000.0,
			max:   500000.0,
			per:   10,
			level: "150,001-500,000",
		},
		{
			min:   500000.0,
			max:   1000000.0,
			per:   15,
			level: "500,001-1,000,000",
		},
		{
			min:   1000000.0,
			max:   2000000.0,
			per:   20,
			level: "1,000,001-2,000,000",
		},
		{
			min:   2000000.0,
			max:   math.MaxFloat64,
			per:   35,
			level: "2,000,001 ขึ้นไป",
		},
	}

	allowanceTable := map[string]float64{
		"donation":  100000.0,
		"k-receipt": kDeduct,
	}

	var taxLevels []request.TaxLevel

	var pay float64

	// ลดหย่อนส่วนตัว
	dummyIncome := totalIncome - personalDeduct

	// ลดหย่อนบริจาค หรือ k - receipt
	for _, allowance := range allowances {
		if allowanceTable[allowance.AllowanceType] > allowance.Amount {
			dummyIncome -= allowance.Amount // ถ้าไม่เกินค่า max ของค่าลดหย่อนให้เอาจำนวนที่ใส่เข้ามาหักออก
		} else {
			dummyIncome -= allowanceTable[allowance.AllowanceType] // ถ้าเกินให้เอาค่า ที่อยู่ในตารางมาใส่ซึงคือค่า default หรือ ค่าที่ admin กำหนด
		}
	}

	// คิด tax
	for i := range taxTable {
		var paylevel float64

		// คิดแต่ละชั้นว่าต้องจ่ายเท่าไหร่
		if taxTable[i].max > dummyIncome && taxTable[i].min < dummyIncome {
			paylevel = (dummyIncome-taxTable[i].min) * taxTable[i].per
		} else if taxTable[i].max < dummyIncome {
			paylevel = (taxTable[i].max-taxTable[i].min) * taxTable[i].per
		}

		// tax ที่ต้องจ่าย
		pay += paylevel / 100

		// สร้างตารางขั้นบันได tax
		taxLevels = append(taxLevels, request.TaxLevel{
			Level: taxTable[i].level,
			Tax:   paylevel / 100,
		})
	}

	// เอา tax ที่จ่ายแล้วมาหักออก
	pay -= wht
	pay = math.Ceil(pay)

	// ถ้ามีเงินที่ต้องรับคืน
	var taxRefund float64
	if pay < 0 {
		taxRefund = pay * (-1)
		pay = 0.0
	}

	taxResponse := respone.TaxResponse{
		Tax:       pay,
		TaxLevel:  taxLevels,
		TaxRefund: taxRefund,
	}

	return taxResponse
}

func (s *TaxService) Cal(requestBody request.TaxRequest) (respone.TaxResponse, error) {
	deduct, err := s.TaxRepo.GetTaxData()
	if err != nil {
		return respone.TaxResponse{}, err
	}
	response := calculateTax(float64(requestBody.TotalIncome), float64(deduct[0].Persernal_Deduct), requestBody.WHT, requestBody.Allowances, float64(deduct[0].K_Reciept_Deduct))
	return response, nil
}
