package database

import (
	"errors"
)

func (db *appdbimpl) SetName(id string, newUserName string) error {
	rows, err := db.c.Query(`SELECT nome FROM Utente`)
	if err != nil {
		return err
	}
	defer func() { _ = rows.Close() }()

	var output string
	for rows.Next() {
		err = rows.Scan(&output)
		if err != nil {
			return err
		}

		if output == newUserName {
			return errors.New("Username already used by another person")
		}
	}

	_, err = db.c.Exec(`UPDATE Utente SET nome = ? WHERE id = ?`, newUserName, id)
	if err != nil {
		return err
	}

	return nil
}
