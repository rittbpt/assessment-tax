package response

import request "github.com/rittbpt/assessment-tax/Request"

type TaxResponse struct {
	Tax       float64             `json:"tax"`
	TaxRefund float64             `json:"taxRefund"`
	TaxLevel  []request.TaxLevel  `json:"taxLevel"`
}


type TaxData struct {
	TotalIncome float64 `json:"totalIncome"`
	Tax         float64 `json:"tax"`
	TaxRefund         float64 `json:"taxRefund"`
}

type TaxCSVResponse struct {
	Taxes []TaxData `json:"taxes"`
}

type UpdateP struct {
	PersonalDeduction float64 `json:"personalDeduction"`
}

type UpdateK struct {
	KReceipt float64 `json:"kReceipt"`
}
