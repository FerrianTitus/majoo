package repository

import (
	"github.com/majoo/database"
	"github.com/majoo/models"
)

func GetUserByUsername(username string, password string) (models.Users, error) {
	db, _ := database.Connect()
	query := GetUserByUsernameQuery
	args := []interface{}{username, password}

	rows, err := db.Query(query, args...)
	if err != nil {
		return models.Users{}, err
	}

	defer rows.Close()
	result := models.Users{}
	for rows.Next() {
		err = rows.Scan(
			&result.ID,
			&result.Name,
			&result.Username,
			&result.Password,
			&result.CreatedAt,
			&result.CreatedBy,
			&result.UpdatedAt,
			&result.UpdatedBy,
		)
	}

	return result, nil
}
