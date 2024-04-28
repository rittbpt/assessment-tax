package service

import (
	"testing"

	"github.com/rittbpt/assessment-tax/Request"
)

func TestCalculateTax(t *testing.T) {
	
	TaxTestTable := []struct {
		totalIncome        float64
		personalDeduct     float64
		wht                float64
		allowances         []request.Allowance
		kDeduct            float64
		expectedTax        float64
		expectedTaxRefund  float64
		expectedTaxLevels []request.TaxLevel
	}{
		{
			totalIncome:       0.0,
			wht:               0.0,
			allowances: []request.Allowance{
			},
			kDeduct:           50000.0,
			personalDeduct:    60000.0,
			expectedTax:       0.0,
			expectedTaxRefund: 0.0,
			expectedTaxLevels: []request.TaxLevel{
				{Level: "0-150,000", Tax: 0.0},
				{Level: "150,001-500,000", Tax: 0.0},
				{Level: "500,001-1,000,000", Tax: 0.0},
				{Level: "1,000,001-2,000,000", Tax: 0.0},
				{Level: "2,000,001 ขึ้นไป", Tax: 0.0},
			},
		},
		{
			totalIncome:       100000.0,
			wht:               0.0,
			allowances: []request.Allowance{
			},
			kDeduct:           50000.0,
			personalDeduct:    60000.0,
			expectedTax:       0.0,
			expectedTaxRefund: 0.0,
			expectedTaxLevels: []request.TaxLevel{
				{Level: "0-150,000", Tax: 0.0},
				{Level: "150,001-500,000", Tax: 0.0},
				{Level: "500,001-1,000,000", Tax: 0.0},
				{Level: "1,000,001-2,000,000", Tax: 0.0},
				{Level: "2,000,001 ขึ้นไป", Tax: 0.0},
			},
		},
		{
			totalIncome:       300000.0,
			wht:               0.0,
			allowances: []request.Allowance{
			},
			kDeduct:           50000.0,
			personalDeduct:    60000.0,
			expectedTax:       9000.0,
			expectedTaxRefund: 0.0,
			expectedTaxLevels: []request.TaxLevel{
				{Level: "0-150,000", Tax: 0.0},
				{Level: "150,001-500,000", Tax: 9000.0},
				{Level: "500,001-1,000,000", Tax: 0.0},
				{Level: "1,000,001-2,000,000", Tax: 0.0},
				{Level: "2,000,001 ขึ้นไป", Tax: 0.0},
			},
		},
		{
			totalIncome:       2160001.0,
			wht:               200000.35,
			allowances: []request.Allowance{
				{AllowanceType: "donation", Amount: 2000000.0},
			},
			kDeduct:           50000.0,
			personalDeduct:    60000.0,
			expectedTax:       110000.0,
			expectedTaxRefund: 0.0,
			expectedTaxLevels: []request.TaxLevel{
				{Level: "0-150,000", Tax: 0.0},
				{Level: "150,001-500,000", Tax: 35000.0},
				{Level: "500,001-1,000,000", Tax: 75000.0},
				{Level: "1,000,001-2,000,000", Tax: 200000.0},
				{Level: "2,000,001 ขึ้นไป", Tax: 0.35},
			},
		},
		{
			totalIncome:       500000.0,
			wht:               0.0,
			allowances: []request.Allowance{
			},
			kDeduct:           50000.0,
			personalDeduct:    60000.0,
			expectedTax:       29000.0,
			expectedTaxRefund: 0.0,
			expectedTaxLevels: []request.TaxLevel{
				{Level: "0-150,000", Tax: 0.0},
				{Level: "150,001-500,000", Tax: 29000.0},
				{Level: "500,001-1,000,000", Tax: 0.0},
				{Level: "1,000,001-2,000,000", Tax: 0.0},
				{Level: "2,000,001 ขึ้นไป", Tax: 0.0},
			},
		},
		{
			totalIncome:       500000.0,
			wht:               0.0,
			allowances: []request.Allowance{
				{AllowanceType: "donation", Amount: 100000.0},
				{AllowanceType: "k-receipt", Amount: 200000.0},
			},
			kDeduct:           50000.0,
			personalDeduct:    60000.0,
			expectedTax:       14000.0,
			expectedTaxRefund: 0.0,
			expectedTaxLevels: []request.TaxLevel{
				{Level: "0-150,000", Tax: 0.0},
				{Level: "150,001-500,000", Tax: 14000.0},
				{Level: "500,001-1,000,000", Tax: 0.0},
				{Level: "1,000,001-2,000,000", Tax: 0.0},
				{Level: "2,000,001 ขึ้นไป", Tax: 0.0},
			},
		},
		{
			totalIncome:       500000.0,
			wht:               0.0,
			allowances: []request.Allowance{
				{AllowanceType: "donation", Amount: 5000.0},
			},
			kDeduct:           50000.0,
			personalDeduct:    60000.0,
			expectedTax:       28500.0,
			expectedTaxRefund: 0.0,
			expectedTaxLevels: []request.TaxLevel{
				{Level: "0-150,000", Tax: 0.0},
				{Level: "150,001-500,000", Tax: 28500.0},
				{Level: "500,001-1,000,000", Tax: 0.0},
				{Level: "1,000,001-2,000,000", Tax: 0.0},
				{Level: "2,000,001 ขึ้นไป", Tax: 0.0},
			},
		},
		{
			totalIncome:       500000.0,
			wht:               20000.0,
			allowances: []request.Allowance{
				{AllowanceType: "donation", Amount: 100000.0},
				{AllowanceType: "k-receipt", Amount: 200000.0},
			},
			kDeduct:           50000.0,
			personalDeduct:    60000.0,
			expectedTax:       0.0,
			expectedTaxRefund: 6000.0,
			expectedTaxLevels: []request.TaxLevel{
				{Level: "0-150,000", Tax: 0.0},
				{Level: "150,001-500,000", Tax: 14000.0},
				{Level: "500,001-1,000,000", Tax: 0.0},
				{Level: "1,000,001-2,000,000", Tax: 0.0},
				{Level: "2,000,001 ขึ้นไป", Tax: 0.0},
			},
		},
	}

	for _, tc := range TaxTestTable {
		taxResponse := CalculateTax(tc.totalIncome, tc.personalDeduct, tc.wht, tc.allowances, tc.kDeduct)

		if taxResponse.Tax != tc.expectedTax {
			t.Errorf("Expected tax: %f, got: %f", tc.expectedTax, taxResponse.Tax)
		}

		if taxResponse.TaxRefund != tc.expectedTaxRefund {
			t.Errorf("Expected tax refund: %f, got: %f", tc.expectedTaxRefund, taxResponse.TaxRefund)
		}

		if len(taxResponse.TaxLevel) != len(tc.expectedTaxLevels) {
			t.Errorf("Expected %d tax levels, got %d", len(tc.expectedTaxLevels), len(taxResponse.TaxLevel))
		}

		for i, expectedLevel := range tc.expectedTaxLevels {
			if taxResponse.TaxLevel[i].Level != expectedLevel.Level || taxResponse.TaxLevel[i].Tax != expectedLevel.Tax {
				t.Errorf("Expected tax level %s: %f, got: %s: %f", expectedLevel.Level, expectedLevel.Tax, taxResponse.TaxLevel[i].Level, taxResponse.TaxLevel[i].Tax)
			}
		}
	}
}
