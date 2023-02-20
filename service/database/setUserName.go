package database

import (
	"errors"
)

var ErrChangeName = errors.New("Username already used by another person")

func (db *appdbimpl) SetName(id string, newUserName string) error {

	// Collect info about the name registred in the database
	rows, err := db.c.Query(`SELECT nome FROM Utente`)
	if err != nil {
		return err
	}

	// Check if the new username is already into the database
	var output string
	for rows.Next() {
		err = rows.Scan(&output)
		if err != nil {
			return err
		}

		if output == newUserName {
			rows.Close()
			return ErrChangeName
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	// Put the information into the table
	rows.Close() // -> Required, otherwise the database remains locked
	_, err = db.c.Exec(`UPDATE Utente SET nome = ? WHERE id = ?`, newUserName, id)
	if err != nil {
		return err
	}

	return nil
}
