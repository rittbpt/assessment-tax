package respone

import "github.com/rittbpt/assessment-tax/Request"

type TaxResponse struct {
	Tax       float64             `json:"tax"`
	TaxRefund float64             `json:"taxRefund"`
	TaxLevel  []request.TaxLevel  `json:"taxLevel"`
}
