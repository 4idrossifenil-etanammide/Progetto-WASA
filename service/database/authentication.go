package database

import (
	"errors"
	"fmt"
)

func (db *appdbimpl) CheckToken(id string, username string) error {

	rows, err := db.c.Query(`SELECT nome FROM Utente WHERE id = ?`, id)
	defer rows.Close()
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
		errLog := fmt.Sprintf("L'ID non corrisponde con l'username. ID: %s, username: %s", id, output)
		return errors.New(errLog)
	}

	return nil
}
