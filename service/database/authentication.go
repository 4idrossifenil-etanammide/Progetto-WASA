package database

import (
	"errors"
)

var ErrAuthentication = errors.New("L'ID non corrisponde con l'username")

func (db *appdbimpl) CheckToken(id string, username string) error {

	rows, err := db.c.Query(`SELECT Nome FROM Utente WHERE id = ?;`, id)
	if err != nil {
		return err
	}

	var output string
	if rows.Next() {
		err = rows.Scan(&output)
		if err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	if username != output {
		return ErrAuthentication
	}

	rows.Close()

	return nil
}
