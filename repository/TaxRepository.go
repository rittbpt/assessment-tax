package repository

import (
    "database/sql"
    "github.com/rittbpt/assessment-tax/model"
)

type TaxRepository struct {
    DB *sql.DB
}

func NewTaxRepository(db *sql.DB) *TaxRepository {
    return &TaxRepository{
        DB: db,
    }
}

func (r *TaxRepository) GetTaxData() ([]model.Tax, error) {
    query := "SELECT k_reciept_deduct, personal_deduct FROM tax WHERE active is true"

    rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var taxData []model.Tax
	for rows.Next() {
		var tax model.Tax
		err := rows.Scan(&tax.K_Reciept_Deduct, &tax.Persernal_Deduct)
		if err != nil {
			return nil, err
		}
		taxData = append(taxData, tax)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	
	return taxData, nil
}
