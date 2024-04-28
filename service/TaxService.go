package service

import (
	"github.com/rittbpt/assessment-tax/Request"
	"github.com/rittbpt/assessment-tax/Response"
	"github.com/rittbpt/assessment-tax/repository"
	"errors"
)

type TaxService struct {
	TaxRepo *repository.TaxRepository
}

func NewTaxService(repo *repository.TaxRepository) *TaxService {
	return &TaxService{
		TaxRepo: repo,
	}
}


func (s *TaxService) ChangeDp(amount float64) (response.UpdateP, error) {
    if amount > 100000 {
        return response.UpdateP{}, errors.New("ค่าลดหย่อนส่วนตัวห้ามเกิน 100,000")
    }

    if amount < 10000 {
        return response.UpdateP{}, errors.New("ค่าลดหย่อนส่วนตัวต้องมีค่ามากกว่า 10,000")
    }

    err := s.TaxRepo.ChangeDp(amount)
    if err != nil {
        return response.UpdateP{}, err
    }

    result := response.UpdateP{
        PersonalDeduction: amount,
    }
    return result, nil
}


func (s *TaxService) ChangeDk(amount float64) (response.UpdateK, error) {
	if amount > 100000 {
		return response.UpdateK{}, errors.New("ค่าลดหย่อน k-receipt ห้ามเกิน 100,000")
	}

	if amount < 0 {
		return response.UpdateK{}, errors.New("ค่าลดหย่อน k-receipt ต้องมีค่ามากกว่า 0")
	}

	err := s.TaxRepo.ChangeDk(amount)
	if err != nil {
		return response.UpdateK{}, err
	}
	result := response.UpdateK{
		KReceipt: amount,
	}
	return result, nil
}



func (s *TaxService) Cal(requestBody request.TaxRequest) (response.TaxResponse, error) {
	deduct, err := s.TaxRepo.GetTaxData()
	if err != nil {
		return response.TaxResponse{}, err
	}
	response := CalculateTax(float64(requestBody.TotalIncome), float64(deduct[0].Persernal_Deduct), requestBody.WHT, requestBody.Allowances, float64(deduct[0].K_Reciept_Deduct))
	return response, nil
}

func (s *TaxService) CalCSV(requestBody []request.TaxRequest) ([]response.TaxData, error) {
	deduct, err := s.TaxRepo.GetTaxData()
	if err != nil {
		return nil, err
	}

	result := make([]response.TaxData, len(requestBody))

	for i := range requestBody {
		taxResponse := CalculateTax(float64(requestBody[i].TotalIncome), float64(deduct[0].Persernal_Deduct), requestBody[i].WHT, requestBody[i].Allowances, float64(deduct[0].K_Reciept_Deduct))

		result[i] = response.TaxData{
			TotalIncome: requestBody[i].TotalIncome,
			Tax:         taxResponse.Tax,
			TaxRefund:   taxResponse.TaxRefund,
		}
	}
	return result, nil
}
