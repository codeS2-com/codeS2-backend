package model

import (
  "fmt"
  "database/sql"
)

type Code struct {
	ID   int    `json:"id"`
	Title string `json:"title"`
}

func (u *Code) GetCode(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT title FROM code WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Title)
}

func FindAll(db *sql.DB) ([]Code, error) {
	statement := fmt.Sprintf("SELECT id, title FROM code")
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Codes := []Code{}

	for rows.Next() {
		var u Code
		if err := rows.Scan(&u.ID, &u.Title); err != nil {
			return nil, err
		}
		Codes = append(Codes, u)
	}

	return Codes, nil
}