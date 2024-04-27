package service

import (
	_ "log"
	"math"
	"github.com/rittbpt/assessment-tax/repository"
	"github.com/rittbpt/assessment-tax/Request"
)

type TaxService struct {
	TaxRepo *repository.TaxRepository
}


func calculations(totalIncome float64, personalDeduct float64) float64 {
	taxTable := []struct {
		min float64
		max float64
		per float64
	}{
		{
			min: 0.0,
			max: 150000.0,
			per: 0.0,
		},
		{
			min: 150001.0,
			max: 500000.0,
			per: 0.1,
		},
		{
			min: 500001.0,
			max: 1000000.0,
			per: 0.15,
		},
		{
			min: 1000001.0,
			max: 2000000.0,
			per: 0.20,
		},
		{
			min: 2000001.0,
			max: math.MaxFloat64,
			per: 0.35,
		},
	}

	var pay float64
	dummyIncome := totalIncome - personalDeduct

	for i := len(taxTable) - 1; i >= 0; i-- {
		if taxTable[i].max >= dummyIncome && taxTable[i].min <= dummyIncome {
			pay += float64(dummyIncome-taxTable[i].min+1) * taxTable[i].per
			dummyIncome -= dummyIncome - taxTable[i].min + 1
		} else if taxTable[i].max < dummyIncome {
			pay += float64(taxTable[i].max-taxTable[i].min - 1) * taxTable[i].per
		}
	}
	
	return pay
}

func NewTaxService(repo *repository.TaxRepository) *TaxService {
	return &TaxService{
		TaxRepo: repo,
	}
}

func (s *TaxService) Cal(requestBody request.TaxRequest) (float64, error) {
    deduct, err := s.TaxRepo.GetTaxData()
    if err != nil {
        return 0.0, err
    }
    pay := calculations(float64(requestBody.TotalIncome), float64(deduct[0].Persernal_Deduct))
    return pay, nil
}
