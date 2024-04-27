package request

type TaxRequest struct {
	TotalIncome  float64 `json:"totalIncome"`
	WHT          float64 `json:"wht"`
	Allowances   []struct {
		AllowanceType string  `json:"allowanceType"`
		Amount        float64 `json:"amount"`
	} `json:"allowances"`
}