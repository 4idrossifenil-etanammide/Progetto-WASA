package database

import (
	"errors"
)

var ErrAuthentication = errors.New("L'ID non corrisponde con l'username")

func (db *appdbimpl) CheckToken(id string, username string) error {

	rows, err := db.c.Query(`SELECT nome FROM Utente WHERE id = ?`, id)
	if err != nil {
		return err
	}
	defer func() { _ = rows.Close() }()

	var output string
	if rows.Next() {
		err = rows.Scan(&output)
		if err != nil {
			return err
		}
	}

	if username != output {
		return ErrAuthentication
	}

	return nil
}
