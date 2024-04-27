package repository

import (
    "database/sql"
    "github.com/rittbpt/assessment-tax/model"
	"context"
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

func (r *TaxRepository) ChangeDp(personalDeduct float64) (string, error) {
	// ดึงค่า k-receipt ล่าสุดมา
	queryLastData := "SELECT k_reciept_deduct FROM tax WHERE active is true"
	var currentKReceiptDeduct float64
	err := r.DB.QueryRow(queryLastData).Scan(&currentKReceiptDeduct)
	if err != nil {
		return "", err
	}

	// เปลี่่ยนค่าที่ active เป็น false
	queryUpdate := "UPDATE tax SET active = false"
	_, err = r.DB.Exec(queryUpdate)
	if err != nil {
		return "", err
	}

	// ใส่ data ใหม่ที่ active ลงไป
	ctx := context.Background()
	queryInsert := "INSERT INTO tax (k_reciept_deduct, personal_deduct , active ,create_time) VALUES ($1, $2 , true , NOW())"
	_, err = r.DB.ExecContext(ctx, queryInsert, currentKReceiptDeduct, personalDeduct)

	if err != nil {
		return "", err
	}

	return "update success", nil
}
