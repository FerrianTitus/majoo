package repository

import (
	"github.com/majoo/database"
	"github.com/majoo/models"
)

func GetOutletByUserId(userId int) ([]models.GetOutletByUserId, error) {
	db, _ := database.Connect()
	query := GetOutletByUserIdQuery
	args := []interface{}{userId}

	rows, err := db.Query(query, args...)
	if err != nil {
		return []models.GetOutletByUserId{}, err
	}

	defer rows.Close()
	var result []models.GetOutletByUserId
	for rows.Next() {
		data := models.GetOutletByUserId{}
		err = rows.Scan(
			&data.OutletName,
			&data.MerchantName,
			&data.UserName,
		)

		result = append(result, data)
	}

	return result, nil
}
