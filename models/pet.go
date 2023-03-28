package models

import (
	"database/sql"
	"fmt"
)

type Pet struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
	City        string `json:"city"`
}

func GetAllPets(db *sql.DB) ([]Pet, error) {
	pets := []Pet{}
	rows, err := db.Query("SELECT id, name, age, description, city FROM pets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if len(pets) == 0 {
		return nil, fmt.Errorf("não foram encontrados dados")
	}

	for rows.Next() {
		pet := Pet{}
		err := rows.Scan(&pet.ID, &pet.Name, &pet.Age, &pet.Description, &pet.City)
		if err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pets, nil
}
