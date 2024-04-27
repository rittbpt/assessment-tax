package request


type TaxRequest struct {
	TotalIncome float64    `json:"totalIncome"`
	WHT         float64    `json:"wht"`
	Allowances  []Allowance `json:"allowances"`
}

type Allowance struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

type TaxLevel struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}

type PersernalDeduct struct {
	Amount float64  `json:"amount"`
}
type K_recieptDeduct struct {
	Amount float64  `json:"amount"`
}